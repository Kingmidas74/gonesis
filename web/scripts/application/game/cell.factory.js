import {Wall, Agent, Empty, Colors } from "../render/index.js";
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
        return initialColor
        let [_, hue, saturation, lightness, alpha] = initialColor.match(/hsla\((\d+),\s*([\d.]+)%,\s*([\d.]+)%,\s*([\d.]+)\)/i);
        lightness = Math.min(Number(lightness) + light, 100);
        return `hsla(${hue}, ${saturation}%, ${lightness}%, ${alpha})`;
    }

    #changeAlpha(initialColor, newAlpha) {
        return initialColor
        // Extract HSLA values using regular expressions
        let [_, hue, saturation, lightness] = initialColor.match(/hsla\((\d+),\s*([\d.]+)%,\s*([\d.]+)%,\s*([\d.]+)\)/i);
        return `hsla(${hue}, ${saturation}%, ${lightness}%, ${newAlpha})`;
    }

    createWall(x, y) {
        return new Wall(x, y, this.#configuration.getInstance().MazeColor);
    }

    createEmpty(x, y, energyPercent) {
        return new Empty(x, y, Colors.YELLOW);
    }

    createAgent(x, y, energy, agentType) {
        let agentColor = null;
        switch (agentType) {
            case AgentType.CARNIVORE:
                agentColor = this.#configuration.getInstance().CarnivoreConfiguration.Color
                break;
            case AgentType.HERBIVORE:
                agentColor = this.#configuration.getInstance().HerbivoreConfiguration.Color
                break;
            case AgentType.DECOMPOSER:
                agentColor = this.#configuration.getInstance().DecomposerConfiguration.Color
                break;
            case AgentType.PLANT:
                agentColor = this.#configuration.getInstance().PlantConfiguration.Color
                break;
            case AgentType.OMNIVORE:
                agentColor = this.#configuration.getInstance().OmnivoreConfiguration.Color
                break;
            default:
                throw "Unknown agent type: " + agentType;
        }

        return new Agent(x, y, agentColor, energy);
    }
}

export { CellFactory };