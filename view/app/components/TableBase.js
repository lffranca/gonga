import * as R from 'ramda'
import {Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow} from "@mui/material";
import React from "react";

function TableBase({rows}) {
    const headers = R.keys(rows[0])
    const values = rows.map(i => R.values(i))

    return (
        <TableContainer component={Paper}>
            <Table sx={{minWidth: 650}} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        {headers.map(i => (<TableCell>{i}</TableCell>))}
                    </TableRow>
                </TableHead>
                <TableBody>
                    {values.map(row => (
                        <TableRow
                            key={row[0]}
                            sx={{'&:last-child td, &:last-child th': {border: 0}}}
                        >
                            {row.map(value => (<TableCell>{value}</TableCell>))}
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    )
}

export default TableBase
