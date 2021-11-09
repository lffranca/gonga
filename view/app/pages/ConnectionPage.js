import React, {useContext, useState} from 'react'
import BasePage from "./BasePage";
import {Button, TableCell, TableRow} from "@mui/material";
import TableBase from "../components/TableBase";
import {GatewayContext} from "../state/GatewayContext";

function ConnectionPage() {
    const {gateway, setGateway} = useContext(GatewayContext)
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
                    name: "Connections",
                }
            ]}
            title="Connections"
            subtitle="Create connections to Kong Nodes and activate the one you want use."
            results={results}
            setResults={setResults}
        >
            <TableBase
                entity="gateway"
                results={results}
                tableHeadComponent={(
                    <TableRow>
                        <TableCell>Name</TableCell>
                        <TableCell>Host</TableCell>
                        <TableCell></TableCell>
                    </TableRow>
                )}
                tableItemsMap={(item, index) => {
                    return (
                        <TableRow
                            key={`table-gateway-item-${index}`}
                            sx={{'&:last-child td, &:last-child th': {border: 0}}}
                        >
                            <TableCell>{item.name}</TableCell>
                            <TableCell>{item.host}</TableCell>
                            <TableCell>
                                {gateway && item && gateway.id === item.id ? (
                                    <Button
                                        variant="contained"
                                        onClick={() => setGateway(null)}
                                    >
                                        Deactivate
                                    </Button>
                                ) : (
                                    <Button
                                        variant="text"
                                        onClick={() => setGateway(item)}
                                    >
                                        Activate
                                    </Button>)}
                            </TableCell>
                        </TableRow>
                    )
                }}
            />
        </BasePage>
    )
}

export default ConnectionPage
