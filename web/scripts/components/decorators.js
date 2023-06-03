import {TemplateParserService} from "./template-parser-service.js";

export function Component({
                              templatePath,
                              stylePath,
                              windowProvider = window,
                              documentProvider = document,
                              logger = console,
                              templateParser = new TemplateParserService(),
                          }) {
    return (OriginalClass) => {
        OriginalClass.stylePath = stylePath;
        OriginalClass.templatePath = templatePath;
        OriginalClass.windowProvider = windowProvider;
        OriginalClass.documentProvider = documentProvider;
        OriginalClass.logger = logger;
        OriginalClass.templateParser = templateParser;
        return OriginalClass;
    };
}