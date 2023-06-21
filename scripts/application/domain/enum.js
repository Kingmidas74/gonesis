const CellType = Object.freeze({
    EMPTY:  "empty",
    WALL: "wall",
    AGENT: "agent",
});

const AgentType = Object.freeze({
    CARNIVORE:  "carnivore",
    HERBIVORE: "herbivore",
    DECOMPOSER: "decomposer",
    PLANT: "plant",
    OMNIVORE: "omnivore",
});

export { CellType, AgentType };