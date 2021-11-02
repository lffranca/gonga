import React, {useState} from 'react'
import BasePage from "./BasePage";
import TableBase from "../components/TableBase";
import {TableCell, TableRow} from "@mui/material";

function ServicePage() {
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
                    name: "Services"
                }
            ]}
            title="Services"
            subtitle="Service entities, as the name implies, are abstractions of each of your own upstream services.
            Examples of Services would be a data transformation microservice, a billing API, etc."
            results={results}
            setResults={setResults}
        >
            <TableBase
                entity="service"
                results={results}
                tableHeadComponent={(
                    <TableRow>
                        <TableCell>Name</TableCell>
                        <TableCell>Host</TableCell>
                        <TableCell>Tags</TableCell>
                        <TableCell>Created At</TableCell>
                    </TableRow>
                )}
                tableItemsMap={(item, index) => {
                    return (
                        <TableRow
                            key={`table-service-item-${index}`}
                            sx={{'&:last-child td, &:last-child th': {border: 0}}}
                        >
                            <TableCell>{item.name}</TableCell>
                            <TableCell>{item.host}</TableCell>
                            <TableCell>{item.tags.join(", ")}</TableCell>
                            <TableCell>{item.created_at}</TableCell>
                        </TableRow>
                    )
                }}
            />
        </BasePage>
    )
}


export default ServicePage
