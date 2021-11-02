import React from 'react'
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Container from "@mui/material/Container";
import Copyright from "../components/Copyright";
import PropTypes from 'prop-types'
import {Breadcrumbs, FormControl, InputLabel, MenuItem, Paper, Select, TextField, Typography} from "@mui/material";
import ItemLink from "../components/ItemLink";
import Button from "@mui/material/Button";
import SearchIcon from "@mui/icons-material/Search";
import {styled} from "@mui/material/styles";

function BasePage({children, isList, breadcrumbs, title, subtitle, results, setResults}) {
    const handleChangeResults = (event) => {
        setResults(event.target.value);
    };

    return (
        <Box
            component="main"
            sx={{
                backgroundColor: (theme) =>
                    theme.palette.mode === 'light'
                        ? theme.palette.grey[100]
                        : theme.palette.grey[900],
                flexGrow: 1,
                height: '100vh',
                overflow: 'auto',
            }}
        >
            <Toolbar/>
            <Container maxWidth="lg" sx={{mt: 4, mb: 4}}>
                <Breadcrumbs aria-label="breadcrumb">
                    {breadcrumbs.map((item, index) => {
                        if ('to' in item) {
                            return (
                                <ItemLink
                                    key={`breadcrumbs-${index}`}
                                    underline="hover"
                                    color="inherit"
                                    to={item.to}>
                                    {item.name}
                                </ItemLink>
                            )
                        }

                        return (
                            <Typography
                                key={`breadcrumbs-${index}`}
                                color="text.primary">
                                {item.name}
                            </Typography>
                        )
                    })}
                </Breadcrumbs>
                <br/>
                <HeaderPage>
                    <Typography variant="h4">{title}</Typography>
                    <br/>
                    <Typography variant="body1">
                        {subtitle}
                    </Typography>
                </HeaderPage>
                <br/>
                {isList ? (
                    <ButtonsContainer>
                        <Button variant="contained">Add New</Button>
                        <div style={{"flex": 1}}/>
                        <Box sx={{display: 'flex', alignItems: 'flex-end'}}>
                            <SearchIcon sx={{color: 'action.active', mr: 1, my: 0.5}}/>
                            <TextField id="input-search-services" label="search..." variant="standard"/>
                        </Box>
                        <FormControlResults>
                            <InputLabel id="input-result-select-label">Results</InputLabel>
                            <Select
                                labelId="input-result-select-label"
                                id="input-result-select"
                                value={results}
                                label="Results"
                                onChange={handleChangeResults}
                            >
                                <MenuItem value={5}>5</MenuItem>
                                <MenuItem value={10}>10</MenuItem>
                                <MenuItem value={25}>25</MenuItem>
                                <MenuItem value={50}>50</MenuItem>
                                <MenuItem value={100}>100</MenuItem>
                            </Select>
                        </FormControlResults>
                    </ButtonsContainer>
                ) : null}
                <br/>
                {children}
                <Copyright sx={{pt: 4}}/>
            </Container>
        </Box>
    )
}

const FormControlResults = styled(FormControl)(({theme}) => ({
    marginLeft: theme.spacing(2)
}))

const ButtonsContainer = styled("div")(({theme}) => ({
    display: "flex",
}))

const HeaderPage = styled(Paper)(({theme}) => ({
    padding: theme.spacing(2),
}))

const breadcrumbsProps = PropTypes.shape({
    name: PropTypes.string,
    to: PropTypes.string,
})

BasePage.propTypes = {
    isList: PropTypes.bool.isRequired,
    breadcrumbs: PropTypes.arrayOf(breadcrumbsProps).isRequired,
    title: PropTypes.string.isRequired,
    subtitle: PropTypes.string.isRequired,
    results: PropTypes.number,
    setResults: PropTypes.func,
}

export default BasePage
