import React, {useState} from 'react'
import BasePage from "./BasePage";

function ConsumerPage() {
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
                    name: "Consumers"
                }
            ]}
            title="Consumers"
            subtitle="The Consumer object represents a consumer - or a user - of an API. You can either rely on Kong
            as the primary datastore, or you can map the consumer list with your database to keep consistency between
            Kong and your existing primary datastore."
            results={results}
            setResults={setResults}
        >
            <div>ConsumerPage</div>
        </BasePage>
    )
}

export default ConsumerPage
