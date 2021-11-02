import React from 'react'
import BasePage from "./BasePage";

function SnapshotPage() {
    return (
        <BasePage
            isList={false}
            breadcrumbs={[
                {
                    name: "Dashboard",
                    to: "/"
                },
                {
                    name: "Snapshots",
                }
            ]}
            title="Snapshots"
            subtitle="Take snapshots of currently active nodes.\nAll Services, Routes, APIs, Plugins, Consumers,
            Upstreams and Targetswill be saved and available for later import."
        >
            <div>SnapshotPage</div>
        </BasePage>
    )
}

export default SnapshotPage
