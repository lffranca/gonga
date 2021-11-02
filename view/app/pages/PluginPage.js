import React, {useState} from 'react'
import BasePage from "./BasePage";

function PluginPage() {
    const [results, setResults] = useState(25)

    return (
        <BasePage
            isList={true}
            breadcrumbs={[
                {
                    name: "Dashboard",
                    to: "/"
                },
                {
                    name: "Plugins"
                }
            ]}
            title="Plugins"
            subtitle="A Plugin entity represents a plugin configuration that will be executed during the HTTP
            request/response workflow, and it's how you can add functionalities to APIs that run behind Kong,
            like Authentication or Rate Limiting for example."
            results={results}
            setResults={setResults}
        >
            <div>PluginPage</div>
        </BasePage>
    )
}

export default PluginPage
