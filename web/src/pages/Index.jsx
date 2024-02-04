import { useLoaderData } from "react-router-dom";
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";

function ItemInfo({ title, description }) {
    return (
        <div className="mb-4 grid grid-cols-[25px_1fr] items-start pb-4 last:mb-0 last:pb-0">
            <span className="flex h-2 w-2 translate-y-1 rounded-full bg-sky-500" />
            <div className="flex items-center space-x-2">
                <p className="text-sm font-medium leading-none">{title}</p>
                <p className="text-sm text-muted-foreground">{description}</p>
            </div>
        </div>
    );
}

function Index() {
    const backendStatus = useLoaderData();
    console.log("backendStatus: ", backendStatus);

    return (
        <Card>
            <CardHeader>
                <CardTitle>go-template</CardTitle>
                <CardDescription>这里是项目详情</CardDescription>
            </CardHeader>

            <CardContent>
                <ItemInfo title="项目版本" description={backendStatus.data.version} />
                <ItemInfo
                    title="服务启动时间"
                    description={backendStatus.data.start_time}
                />
            </CardContent>

            <CardFooter>
                <p>Copyright @xpzouying/</p>
            </CardFooter>
        </Card>
    );
}

export default Index;
