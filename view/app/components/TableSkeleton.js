import React from 'react'
import Skeleton from "@mui/material/Skeleton";
import {Paper} from "@mui/material";
import {styled} from "@mui/material/styles";

function TableSkeleton() {
    return (
        <Container>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
            <Skeleton variant="text"/>
        </Container>
    )
}

const Container = styled(Paper)(({theme}) => ({
    padding: theme.spacing(2),
}))

export default TableSkeleton
