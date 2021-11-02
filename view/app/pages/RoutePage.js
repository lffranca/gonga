import React, {useState} from 'react'
import BasePage from "./BasePage";

function RoutePage() {
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
                    name: "Routes"
                }
            ]}
            title="Routes"
            subtitle="The Route entities defines rules to match client requests. Each Route is associated with a Service,
            and a Service may have multiple Routes associated to it. Every request matching a given Route will be proxied
            to its associated Service."
            results={results}
            setResults={setResults}
        >
            <div>RoutePage</div>
        </BasePage>
    )
}

export default RoutePage
