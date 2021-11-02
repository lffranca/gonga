import React from 'react'
import BasePage from "./BasePage";

function InfoPage() {
    return (
        <BasePage
            isList={false}
            breadcrumbs={[
                {
                    name: "Dashboard",
                    to: "/"
                },
                {
                    name: "Information"
                }
            ]}
            title="Node Info"
            subtitle="Generic details about the node."
        >
            <div>InfoPage</div>
        </BasePage>
    )
}

export default InfoPage
