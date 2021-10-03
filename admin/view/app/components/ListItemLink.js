import React, {forwardRef, useMemo} from 'react'
import {Link as RouterLink} from "react-router-dom"
import {ListItem} from "@mui/material";

const ListItemLink = ({to, ...propsToItem}) => {
    const renderLink = useMemo(
        () =>
            forwardRef((itemProps, ref) => (
                <RouterLink to={to} ref={ref} {...itemProps} />
            )),
        [to]
    );

    return (
        <ListItem component={renderLink} {...propsToItem}>
            {propsToItem.children}
        </ListItem>
    );
};

export default ListItemLink;
