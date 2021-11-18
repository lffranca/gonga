import {
    Button,
    CircularProgress,
    Paper,
    Snackbar,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow
} from "@mui/material";
import React, {useContext, useEffect, useState} from "react";
import {listItems} from "../providers/list";
import TableSkeleton from "./TableSkeleton";
import PropTypes from "prop-types";
import IconButton from "@mui/material/IconButton";
import CloseIcon from "@mui/icons-material/Close";
import {GatewayContext} from "../state/GatewayContext";

function TableBase({entity, results, tableHeadComponent, tableItemsMap}) {
    const {gateway} = useContext(GatewayContext)

    const [items, setItems] = useState([])
    const [offset, setOffset] = useState("")
    const [loading, setLoading] = useState(true)
    const [loadingMore, setLoadingMore] = useState(false)
    const [message, setMessage] = useState("")

    async function getItems(offsetParam = "") {
        if (offsetParam === "") {
            setLoading(true)
        } else {
            setLoadingMore(true)
        }

        try {
            const id = gateway && gateway.id ? gateway.id : ""
            const {data, options} = await listItems(entity, id, results, offsetParam)
            if (offsetParam === "") {
                setItems(data)
            } else {
                setItems([...items, ...data])
            }

            if (options != null) {
                if ('offset' in options) {
                    setOffset(options['offset'])
                }
            } else {
                setOffset("")
            }
        } catch (e) {
            console.error("listItems: ", e)
            setMessage(e.message)
        }

        setLoading(false)
        setLoadingMore(false)
    }

    useEffect(() => {
        getItems()
            .then(() => {
            })
    }, [results])


    return (
        <>
            <Snackbar
                anchorOrigin={{
                    vertical: 'bottom',
                    horizontal: 'right',
                }}
                open={message.length > 0}
                autoHideDuration={6000}
                onClose={() => setMessage("")}
                message={message}
                action={
                    <IconButton
                        aria-label="close"
                        color="inherit"
                        sx={{p: 0.5}}
                        onClick={() => setMessage("")}
                    >
                        <CloseIcon/>
                    </IconButton>
                }
            />
            {loading ? (
                <TableSkeleton/>
            ) : (
                <TableContainer component={Paper}>
                    <Table sx={{minWidth: 650}} aria-label="list of the services">
                        <TableHead>
                            {tableHeadComponent}
                        </TableHead>
                        <TableBody>
                            {items ? items.map(tableItemsMap) : null}
                            {offset ? (
                                <TableRow
                                    key={`table-service-item-more`}
                                    sx={{'&:last-child td, &:last-child th': {border: 0}}}
                                >
                                    <TableCell variant="footer" align="center" colSpan={4}>
                                        {loadingMore ? (
                                            <CircularProgress/>
                                        ) : (
                                            <Button onClick={() => getItems(offset)} variant="text">See more</Button>
                                        )}
                                    </TableCell>
                                </TableRow>
                            ) : null}
                        </TableBody>
                    </Table>
                </TableContainer>
            )}
        </>
    )
}

TableBase.propTypes = {
    entity: PropTypes.string.isRequired,
    results: PropTypes.number.isRequired,
    tableHeadComponent: PropTypes.element.isRequired,
    tableItemsMap: PropTypes.func.isRequired,
}

export default TableBase
