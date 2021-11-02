import React from 'react'
import BasePage from "./BasePage";

function ConnectionPage() {
    return (
        <BasePage
            isList={false}
            breadcrumbs={[
                {
                    name: "Dashboard",
                    to: "/"
                },
                {
                    name: "Connections",
                }
            ]}
            title="Connections"
            subtitle="Create connections to Kong Nodes and activate the one you want use."
        >
            <div>ConnectionPage</div>
        </BasePage>
    )
}

export default ConnectionPage
