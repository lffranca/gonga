import React from 'react'
import Overview from "../components/Overview";
import BasePage from "./BasePage";

function DashboardPage() {
    return (
        <BasePage
            isList={false}
            breadcrumbs={[
                {
                    name: "Dashboard"
                }
            ]}
            title="Overview"
            subtitle="All details of the nodes."
        >
            <Overview/>
        </BasePage>
    )
}

export default DashboardPage
