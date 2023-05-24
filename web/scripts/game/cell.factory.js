import {Wall, Agent, Empty, Organic} from "../render/cell.js";
import {AgentType} from "../engine/domain.js";

class CellFactory {
    /**
     * @type {ConfigurationProvider} The configuration of the game.
     */
    #configuration;

    #cellColors;

    /**
     * Constructs a new instance of CellFactory.
     * @param {ConfigurationProvider} configurationProvider
     */
    constructor(configurationProvider) {
        this.#configuration = configurationProvider;
    }

    #calculateColor(initialColor, light) {
        // Extract HSL values using regular expressions
        let [_, hue, saturation, lightness, alpha] = initialColor.match(/hsla\((\d+),\s*([\d.]+)%,\s*([\d.]+)%,\s*([\d.]+)\)/i);

        // Convert the lightness to a number
        lightness = Number(lightness);

        // Increase the lightness by the amount specified by `light`, ensuring it doesn't exceed 100
        lightness = Math.min(lightness + light, 100);

        // Construct the new color
        let newColor = `hsla(${hue}, ${saturation}%, ${lightness}%, ${alpha})`;

        return newColor;
    }

    #changeAlpha(initialColor, newAlpha) {
        // Extract HSLA values using regular expressions
        let [_, hue, saturation, lightness] = initialColor.match(/hsla\((\d+),\s*([\d.]+)%,\s*([\d.]+)%,\s*([\d.]+)\)/i);

        // Construct the new color with the new alpha value
        let newColor = `hsla(${hue}, ${saturation}%, ${lightness}%, ${newAlpha})`;

        return newColor;
    }

    createWall(x, y) {
        return new Wall(x, y, this.#configuration.getInstance().MazeColor);
    }

    createEmpty(x, y, energyPercent) {
        return new Empty(x, y, this.#changeAlpha("hsla(60, 100%, 50%, 1.0)", energyPercent));
    }

    createAgent(x, y, energy, agentType) {
        let agentColor = this.#configuration.getInstance().AgentConfiguration.HerbivoreColor;
        switch (agentType) {
            case AgentType.CARNIVORE:
                agentColor = this.#configuration.getInstance().AgentConfiguration.CarnivoreColor;
                break;
            case AgentType.HERBIVORE:
                agentColor = this.#configuration.getInstance().AgentConfiguration.HerbivoreColor;
                break;
            case AgentType.DECOMPOSER:
                agentColor = this.#configuration.getInstance().AgentConfiguration.DecomposerColor;
                break;
            case AgentType.PLANT:
                agentColor = this.#configuration.getInstance().AgentConfiguration.PlantColor;
                break;
            case AgentType.OMNIVORE:
                agentColor = this.#configuration.getInstance().AgentConfiguration.OmnivoreColor;
                break;
        }

        //let initialColor = this.#changeAlpha(agentColor, .5)
        //const energyPercent = (energy / (this.#configuration.getInstance().AgentConfiguration.MaxEnergy));
        //initialColor = this.#changeAlpha(initialColor, energyPercent);

        return new Agent(x, y, agentColor, energy);
    }

    createOrganic(x, y) {
        return new Organic(x, y, this.#configuration.getInstance().AgentConfiguration.PlantColor);
    }
}

export { CellFactory };