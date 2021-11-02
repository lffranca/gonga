import React from 'react'
import BasePage from "./BasePage";

function UserPage() {
    return (
        <BasePage
            isList={false}
            breadcrumbs={[
                {
                    name: "Dashboard",
                    to: "/"
                },
                {
                    name: "Users",
                }
            ]}
            title="Users"
            subtitle="Information of the users."
        >
            <div>UserPage</div>
        </BasePage>
    )
}

export default UserPage
