import React from 'react'
import BasePage from "./BasePage";

function SettingsPage() {
    return (
        <BasePage
            isList={false}
            breadcrumbs={[
                {
                    name: "Dashboard",
                    to: "/"
                },
                {
                    name: "Settings",
                }
            ]}
            title="Settings"
            subtitle="Settings."
        >
            <div>SettingsPage</div>
        </BasePage>
    )
}

export default SettingsPage
