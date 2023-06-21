export class COLOR_PICKER extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser().catch((err) => {
            COLOR_PICKER.logger.error(err);
        });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            COLOR_PICKER.windowProvider.fetch(
                new URL(COLOR_PICKER.stylePath, new URL(import.meta.url)).href
            ),
            COLOR_PICKER.windowProvider.fetch(
                new URL(COLOR_PICKER.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = COLOR_PICKER.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    /**
     * @param {string} hslaColor - color in hsla format
     */
    set value(hslaColor) {
        if (!this.isConnected) {
            this.#pendingData = hslaColor;
            return;
        }

        this.#template
            .then((templateContent) => {
                const template = COLOR_PICKER.documentProvider.createElement("template");
                template.innerHTML = COLOR_PICKER.templateParser?.parse(templateContent, {
                    color: hslaColor,
                    title: this.getAttribute('data-title') || 'Color',
                });
                this.#shadow.appendChild(template.content.cloneNode(true));

                this.#shadow.querySelector("input").value = this.#hslaToHex(hslaColor);
                this.#shadow.querySelector(".value").textContent = hslaColor;

                this.#shadow.addEventListener('change', (e) => {
                    if (e.target.type !== 'color') {
                        return;
                    }

                    const hslaColor = this.#hexToHSLA(e.target.value);
                    this.#shadow.querySelector(".value").textContent = hslaColor;

                    this.dispatchEvent(new COLOR_PICKER.windowProvider.CustomEvent('change', {
                        detail: { value: hslaColor }
                    }))
                });

            })
            .catch((err) => {
                COLOR_PICKER.logger.error(err);
            });
    }

    get value() {
        return this.#hexToHSLA(this.#shadow.querySelector("input")?.value)
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

    #hslaToHex = (hsla) => {
        const hslaInArray = hsla.substring(5, hsla.length-1).replace(/ /g, '').split(',');

        let h = parseInt(hslaInArray[0]) / 360; // we need to convert it to be between 0 to 1
        let s = parseInt(hslaInArray[1]) / 100; // we need to convert it to be between 0 to 1
        let l = parseInt(hslaInArray[2]) / 100; // we need to convert it to be between 0 to 1
        let a = parseFloat(hslaInArray[3]); // we need to convert it to be between 0 to 1

        let rgba = this.#hslaToRgba(h, s, l, a);
        return this.#rgbaToHex(rgba);
    }

    #hslaToRgba = (h, s, l, a) => {
        let r, g, b;

        if(s === 0){
            r = g = b = l; // achromatic
        } else {
            const hue2rgb = (p, q, t) => {
                if(t < 0) t += 1;
                if(t > 1) t -= 1;
                if(t < 1/6) return p + (q - p) * 6 * t;
                if(t < 1/2) return q;
                if(t < 2/3) return p + (q - p) * (2/3 - t) * 6;
                return p;
            };
            let q = l < 0.5 ? l * (1 + s) : l + s - l * s;
            let p = 2 * l - q;
            r = hue2rgb(p, q, h + 1/3);
            g = hue2rgb(p, q, h);
            b = hue2rgb(p, q, h - 1/3);
        }
        return 'rgba('+Math.round(r * 255)+','+Math.round(g * 255)+','+Math.round(b * 255)+','+a+')';

    }

    #rgbaToHex = (rgba) => {
        const sep = rgba.indexOf(",") > -1 ? "," : " ";
        rgba = rgba.substr(5).split(")")[0].split(sep);

        // Strip the slash if using space-separated syntax
        if (rgba.indexOf("/") > -1)
            rgba.splice(3,1);

        for (let R in rgba) {
            let r = rgba[R];
            if (r.indexOf("%") > -1) {
                const p = r.substr(0,r.length - 1) / 100;

                if (R < 3) {
                    rgba[R] = Math.round(p * 255);
                } else {
                    rgba[R] = p;
                }
            }
        }

        let r = (+rgba[0]).toString(16),
            g = (+rgba[1]).toString(16),
            b = (+rgba[2]).toString(16);

        if (r.length === 1)
            r = "0" + r;
        if (g.length === 1)
            g = "0" + g;
        if (b.length === 1)
            b = "0" + b;

        return "#" + r + g + b;
    }

    #hexToHSLA = (hex, alpha = 1) => {
        // convert hex to RGB first
        let r = 0, g = 0, b = 0;
        if (hex.length === 4) {
            r = "0x" + hex[1] + hex[1];
            g = "0x" + hex[2] + hex[2];
            b = "0x" + hex[3] + hex[3];
        } else if (hex.length === 7) {
            r = "0x" + hex[1] + hex[2];
            g = "0x" + hex[3] + hex[4];
            b = "0x" + hex[5] + hex[6];
        }

        // then to HSL
        r /= 255;
        g /= 255;
        b /= 255;
        let max = Math.max(r, g, b), min = Math.min(r, g, b);
        let h = 0, s = 0, l = (max + min) / 2;

        if (max === min) {
            h = s = 0; // achromatic
        } else {
            let d = max - min;
            s = l > 0.5 ? d / (2 - max - min) : d / (max + min);
            switch(max) {
                case r: h = (g - b) / d + (g < b ? 6 : 0); break;
                case g: h = (b - r) / d + 2; break;
                case b: h = (r - g) / d + 4; break;
            }
            h /= 6;
        }
        return `hsla(${COLOR_PICKER.windowProvider.Math.round(h * 360)}, ${COLOR_PICKER.windowProvider.Math.round(s * 100)}%, ${COLOR_PICKER.windowProvider.Math.round(l * 100)}%, ${alpha})`;
    }
}
