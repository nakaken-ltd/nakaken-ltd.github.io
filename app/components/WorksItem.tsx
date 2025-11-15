import Image, {StaticImageData} from "next/image";
import header from "@/public/header.webp";

export type Paragraph = {
  text: string;
  imgSrc: StaticImageData;
  alt: string;
}

export type Item = {
  title: string;
  paragraphs: Paragraph[];
};

export default function WorksItem({title, paragraphs}: Item) {
  const ParagraphComponent = ({text, imgSrc, alt, index}: Paragraph & { index: number }) => {
    return <div
      className={`
        flex flex-col gap-4 sm:flex-row mb-16
        ${index % 2 === 0 ? "sm:flex-row-reverse" : ""}
      `}
    >
      <Image
        src={imgSrc}
        alt={alt}
        width={1280}
        height={720}
        className="w-full sm:w-[48%] object-cover"
      />
      <p className="w-full sm:text-lg mx-2 text-justify">{text}</p>
    </div>
  };

  return (
    <div>
      <div className="mb-4">
        <Image
          src={header}
          alt={title}
          width={32}
          height={32}
          style={{display: "inline", marginRight: '10px'}}
        />
        <span className="font-bold align-[-6px] text-2xl">{title}</span>
      </div>

      {paragraphs.map((p, i) => (
        <ParagraphComponent key={i} index={i} {...p} />
      ))}
    </div>
  );
}
