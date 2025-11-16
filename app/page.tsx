import Image from "next/image";
import MangaFrames from "./components/MangaFrames"
import MenuBar from "@/app/components/MenuBar";

import jumbotron from "@/public/top_jumbotron.avif";

export default function Home() {
  return (
    <div className="">
      <div className="shadow-sides bg-white max-w-screen-lg mx-auto">
        {/*Menubar blank*/}
        <div className="w-auto h-[64px]"></div>

        <div className="max-w-screen-lg m-auto">
          <Image src={jumbotron} alt="中原建設工業"></Image>
        </div>

        <div className="lg:max-w-screen-lg mx-auto mt-8">
          <MangaFrames/>
        </div>
      </div>
      <MenuBar></MenuBar>
    </div>
  );
}
