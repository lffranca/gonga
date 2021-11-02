import axios from 'axios'
import {API_URI} from "../config/api";

export const provider = axios.create({
    baseURL: API_URI,
    headers: {'Content-Type': 'application/json', "Accept": "*"}
});
