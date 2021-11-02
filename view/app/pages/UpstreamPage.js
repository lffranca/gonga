import React, {useState} from 'react'
import BasePage from "./BasePage";

function UpstreamPage() {
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
                    name: "Upstreams"
                }
            ]}
            title="Upstreams"
            subtitle="The upstream object represents a virtual hostname and can be used to loadbalance incoming
            requests over multiple services (targets). So for example an upstream named service.v1.xyz with an
            API object created with an upstream_url=https://service.v1.xyz/some/path. Requests for this API would
            be proxied to the targets defined within the upstream."
            results={results}
            setResults={setResults}
        >
            <div>UpstreamPage</div>
        </BasePage>
    )
}

export default UpstreamPage
