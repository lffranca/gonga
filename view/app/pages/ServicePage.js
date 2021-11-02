import React from 'react'
import BasePage from "./BasePage";
import {Breadcrumbs, Pagination, Paper, TextField, Typography} from "@mui/material";
import ItemLink from "../components/ItemLink";
import {styled} from "@mui/material/styles";
import TableBase from "../components/TableBase";
import Button from "@mui/material/Button";
import SearchIcon from '@mui/icons-material/Search';
import Box from "@mui/material/Box";

const rows = [
    {
        "connect_timeout": 60000,
        "created_at": "1970-01-19T11:44:54.051-03:00",
        "host": "localhost",
        "id": "1012b142-d158-4a51-af1a-f69bb47449db",
        "name": "Workflow_RS",
        "path": "/",
        "port": 8093,
        "protocol": "http",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": "1970-01-19T19:20:22.741-03:00",
        "write_timeout": 60000,
        "tags": [
            "devportal-dev",
            "apimanagement-dev"
        ]
    },
    {
        "connect_timeout": 60000,
        "created_at": "1970-01-19T11:43:51.874-03:00",
        "host": "localhost",
        "id": "32966392-4fe0-4ccb-b143-16fe62904543",
        "name": "APIManager_RS",
        "path": "/",
        "port": 8092,
        "protocol": "http",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": "1970-01-19T19:20:22.741-03:00",
        "write_timeout": 60000,
        "tags": [
            "devportal-dev",
            "apimanagement-dev"
        ]
    },
    {
        "connect_timeout": 60000,
        "created_at": "1970-01-19T11:43:35.706-03:00",
        "host": "localhost",
        "id": "6e19840f-314c-4234-a1e1-84beea9b7678",
        "name": "HealthCheck_RS",
        "path": "/",
        "port": 8094,
        "protocol": "http",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": "1970-01-19T19:20:22.741-03:00",
        "write_timeout": 60000,
        "tags": [
            "devportal-dev",
            "apimanagement-dev"
        ]
    },
    {
        "connect_timeout": 60000,
        "created_at": "1970-01-19T19:13:11.392-03:00",
        "host": "pokemon-api",
        "id": "705a311c-c94d-4410-840c-44f44d9bc6c1",
        "name": "pokemon-api",
        "path": "/api/v1",
        "port": 80,
        "protocol": "https",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": "1970-01-19T19:13:11.392-03:00",
        "write_timeout": 60000,
        "tags": [
            "apimanagement-via",
            "template-docker"
        ]
    },
    {
        "connect_timeout": 60000,
        "created_at": "1970-01-19T13:31:58.091-03:00",
        "host": "localhost",
        "id": "784dcec0-e146-492a-8605-4042d252130f",
        "name": "APIManagementProduct_RS",
        "path": "/",
        "port": 8098,
        "protocol": "http",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": "1970-01-19T19:20:22.741-03:00",
        "write_timeout": 60000,
        "tags": [
            "devportal-dev",
            "apimanagement-dev"
        ]
    },
    {
        "connect_timeout": 60000,
        "created_at": "1970-01-19T11:43:53.119-03:00",
        "host": "localhost",
        "id": "833d9cab-f4b0-4f30-99b3-d43d3dd45394",
        "name": "AccountService_RS",
        "path": "/",
        "port": 8096,
        "protocol": "http",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": "1970-01-19T19:20:22.741-03:00",
        "write_timeout": 60000,
        "tags": [
            "devportal-dev",
            "apimanagement-dev"
        ]
    },
    {
        "connect_timeout": 60000,
        "created_at": "1970-01-19T18:38:40.224-03:00",
        "host": "localhost",
        "id": "d6564cb9-b452-483e-8b9d-9e7ce8e0e0c6",
        "name": "gdhuo_RS",
        "path": "/",
        "port": 8081,
        "protocol": "http",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": "1970-01-19T19:20:22.741-03:00",
        "write_timeout": 60000,
        "tags": [
            "devportal-dev",
            "apimanagement-dev"
        ]
    },
    {
        "connect_timeout": 60000,
        "created_at": "1970-01-19T11:43:37.742-03:00",
        "host": "localhost",
        "id": "d80a72ef-d0fb-4601-a484-62245fc519ff",
        "name": "ViaAPI_RS",
        "path": "/",
        "port": 8091,
        "protocol": "http",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": "1970-01-19T19:20:22.741-03:00",
        "write_timeout": 60000,
        "tags": [
            "devportal-dev",
            "apimanagament-dev",
            "viaapi-dev"
        ]
    }
]

function ServicePage() {
    return (
        <BasePage>
            <Breadcrumbs aria-label="breadcrumb">
                <ItemLink underline="hover" color="inherit" to="/">
                    Dashboard
                </ItemLink>
                <Typography color="text.primary">Services</Typography>
            </Breadcrumbs>
            <br/>
            <HeaderPage>
                <Typography variant="h4">Services</Typography>
                <br/>
                <Typography variant="body1">
                    Service entities, as the name implies, are abstractions of each of your own upstream services.
                    Examples of Services would be a data transformation microservice, a billing API, etc.
                </Typography>
            </HeaderPage>
            <br/>
            <ButtonsContainer>
                <Button variant="contained">Add New Service</Button>
                <div style={{"flex": 1}}/>
                <Box sx={{display: 'flex', alignItems: 'flex-end'}}>
                    <SearchIcon sx={{color: 'action.active', mr: 1, my: 0.5}}/>
                    <TextField id="input-search-services" label="search..." variant="standard"/>
                </Box>
            </ButtonsContainer>
            <br/>
            <TableBase rows={rows.map(({name, host, tags, created_at}) => ({
                "Name": name,
                "Host": host,
                "Tags": tags.join(", "),
                "Created At": created_at,
            }))}/>
            <br/>
            <ButtonsContainer>
                <div style={{"flex": 1}}/>
                <Pagination count={10}/>
            </ButtonsContainer>
        </BasePage>
    )
}

const ButtonsContainer = styled("div")(({theme}) => ({
    display: "flex",
}))

const HeaderPage = styled(Paper)(({theme}) => ({
    padding: theme.spacing(2),
}))

export default ServicePage
