import Image from "next/image";
import aboutUs from "@/public/text_about_us.webp";
import works from "@/public/text_works.webp";
import logo from "@/public/logo_b.webp";
import recruit from "@/public/text_recruit.webp";
import inquiry from "@/public/text_inquiry.webp";
import Link from "next/link";

const textHeight = 28;

const aboutUsWidth = aboutUs.width * (textHeight / aboutUs.height);
const worksWidth = works.width * (textHeight / works.height);
const logoWidth = logo.width * (40 / logo.height);
const jobsWidth = recruit.width * (textHeight / recruit.height);
const inquiryWidth = inquiry.width * (textHeight / inquiry.height);

export default function MenuBar() {
  return <div id="menu" className="fixed top-0 w-full">
    <div className="flex items-center justify-center h-[64px] max-w-screen-lg bg-white m-auto">
      <Link href="/about" className="hidden md:block flex-grow relative pl-4 h-full cursor-pointer">
        <Image
          src={aboutUs}
          alt="弊社について"
          width={aboutUsWidth}
          objectPosition="center center"
          style={{margin: 'auto', position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%,-50%)', zIndex: 20}}
        />
      </Link>
      <Link href="/works" className="hidden md:block flex-grow relative h-full cursor-pointer">
          <Image
            src={works}
            alt="工事実績"
            width={worksWidth}
            objectPosition="center center"
            style={{margin: 'auto', position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%,-50%)'}}
          />
      </Link>
      <Link href="/" className="block w-[100px] relative h-full cursor-pointer">
        <Image
          src={logo}
          alt="ロゴ"
          width={logoWidth}
          objectPosition="center center"
          style={{margin: 'auto', position: 'absolute', top: '48%', left: '50%', transform: 'translate(-50%,-50%)'}}
        />
      </Link>
      <Link href="/recruit" className="hidden md:block flex-grow relative h-full cursor-pointer">
        <Image
          src={recruit}
          alt="採用情報"
          width={jobsWidth}
          objectPosition="center center"
          style={{margin: 'auto', position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%,-50%)'}}
        />
      </Link>
      <Link href="/inquiry" className="hidden md:block flex-grow relative pr-4 h-full cursor-pointer">
        <Image
          src={inquiry}
          alt="お問い合わせ"
          width={inquiryWidth}
          objectPosition="center center"
          style={{margin: 'auto', position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%,-50%)'}}
        />
      </Link>
    </div>
  </div>
}