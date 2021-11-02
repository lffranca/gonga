import {provider} from "./provider";

export async function listItems(entity = "", results = 25, offset = "") {
    const path = `/${entity}?size=${results}&offset=${offset}`
    const {data} = await provider.get(path)
    return data
}
