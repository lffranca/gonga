import React, {useState} from 'react'
import BasePage from "./BasePage";
import {TableCell, TableRow} from "@mui/material";
import TableBase from "../components/TableBase";

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
            <TableBase
                entity="route"
                results={results}
                tableHeadComponent={(
                    <TableRow>
                        <TableCell>Name</TableCell>
                        <TableCell>Tags</TableCell>
                        <TableCell>Paths</TableCell>
                        <TableCell>Created At</TableCell>
                    </TableRow>
                )}
                tableItemsMap={(item, index) => {
                    return (
                        <TableRow
                            key={`table-service-item-${index}`}
                            sx={{'&:last-child td, &:last-child th': {border: 0}}}
                        >
                            <TableCell>{item.name ? item.name : item.id}</TableCell>
                            <TableCell>{item.tags ? item.tags.join(",\n") : null}</TableCell>
                            <TableCell>{item.paths ? item.paths.join(",\n") : null}</TableCell>
                            <TableCell>{item.created_at}</TableCell>
                        </TableRow>
                    )
                }}
            />
        </BasePage>
    )
}

export default RoutePage
