export class TemplateParserService {
    constructor() {}

    /**
     *
     * @param {String} template raw string with template literals
     * @param {Object} data Any object with values for template
     */
    parse(template, data= {}) {
        const regexFor =
            /{%\s*for\s+(\w+)\s+in\s+(\w+)\s*%}([\s\S]+?){%\s*endfor\s*%}/gi;
//        const regexIf = /{%\s*if\s+(\w+)\s*%}([\s\S]+?){%\s*endif\s*%}/gi;
        const regexIf = /{%\s*if\s+(\w+)\s*%}([\s\S]+?){% else %}([\s\S]+?){%\s*endif\s*%}/gi;
        let match;
        let output = template;

        // handle 'for' directive
        while ((match = regexFor.exec(template)) !== null) {
            const loopVariable = match[1];
            const loopArray = data[match[2]];
            const loopBlock = match[3];

            let renderedLoop = "";
            loopArray.forEach((item) => {
                const loopData = {};
                Object.assign(loopData, item);
                renderedLoop += this.parse(loopBlock, loopData);
            });
            output = output.replace(match[0], renderedLoop);
        }

        // handle 'if' directive
        while ((match = regexIf.exec(output)) !== null) {
            const condition = match[1];
            const ifBlock = match[2];
            const elseBlock = match[3] || "";

            if (data[condition]) {
                output = output.replace(match[0], this.parse(ifBlock, data));
            } else {
                output = output.replace(match[0], this.parse(elseBlock, data));
            }
        }

        output = output.replace(
            /{{\s*([^}\s]+)\s*}}/gi,
            (_, key) => data[key.split(".").pop()]
        );
        return output;
    }
}
