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

export default function WorksItem({ title, paragraphs }: Item) {
  const ParagraphComponent = ({ text, imgSrc, alt, index }: Paragraph & { index: number }) => {
    return <div
      className={`
        flex flex-col gap-8 md:flex-row mt-8
        ${index % 2 === 0 ? "md:flex-row-reverse" : ""}
      `}
    >
      <p className="w-full md:text-lg mx-2 text-justify">{text}</p>
      <Image
        src={imgSrc}
        alt={alt}
        width={1280}
        height={720}
        className="w-full md:w-[48%] object-cover"
      />
    </div>
  };

  return (
    <div>
      <div className="mt-16">
        <Image
          src={header}
          alt={title}
          width={32}
          height={32}
          style={{display: "inline", marginRight: '10px'}}
        />
        <span className="font-bold">{title}</span>
      </div>

      {paragraphs.map((p, i) => (
        <ParagraphComponent key={i} index={i} {...p} />
      ))}
    </div>
  );
}
