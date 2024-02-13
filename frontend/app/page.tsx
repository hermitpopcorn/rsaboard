import classNames from "classnames";
import PageContent from "./content";
import { Nanum_Pen_Script } from "next/font/google";

const font = Nanum_Pen_Script({
  weight: "400",
  subsets: ["latin"],
});

export default function Page() {
  return (
    <div className="container min-h-screen flex flex-col justify-center items-center">
      <main className="flex flex-col justify-center items-center">
        <h1 className={classNames(font.className, "text-4xl")}>rsaboard</h1>
        <PageContent backendUrl={process.env.BACKEND_URL!} />
      </main>
    </div>
  );
}
