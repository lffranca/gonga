import {provider} from "./provider";

export async function listItems(entity = "", gateway = "", results = 25, offset = "") {
    const path = `/${entity}?gateway=${gateway}&size=${results}&offset=${offset}`
    const {data} = await provider.get(path)
    return data
}
