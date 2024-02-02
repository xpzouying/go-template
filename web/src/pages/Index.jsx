import {
    useLoaderData,
} from "react-router-dom";


function Index() {
    const backendStatus = useLoaderData();
    console.log("backendStatus: ", backendStatus);

    return (
        <div>Index</div>
    )
}

export default Index