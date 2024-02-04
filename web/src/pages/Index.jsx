import { Link, useLoaderData } from "react-router-dom";
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
    <div className="flex items-center justify-between space-x-4">
      <div className="flex items-center space-x-4">
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

      <CardContent className="grid gap-6">
        <ItemInfo title="项目版本" description={backendStatus.data.version} />
        <ItemInfo
          title="启动时间"
          description={backendStatus.data.start_time}
        />
      </CardContent>

      <CardFooter>
        <p className="text-center text-sm text-muted-foreground">
          项目地址：
          <Link
            to="https://github.com/xpzouying/go-template"
            className="underline underline-offset-4 hover:text-primary"
          >
            xpzouying/go-template
          </Link>
        </p>
      </CardFooter>
    </Card>
  );
}

export default Index;
