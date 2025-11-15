'use client'

import Image from "next/image";
import MenuBar from "@/app/components/MenuBar";
import WorksItem, { Item } from "@/app/components/WorksItem";

import works from '../../public/text_works.webp';
import bang from '../../public/bang.webp';
import placeholder from '../../public/works_placeholder.webp'

const items: Item[] = [
  {
    title: "除草工事",
    paragraphs: [
      {
        text: "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。",
        imgSrc: placeholder,
        alt: "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。"
      },
      {
        text: "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。",
        imgSrc: placeholder,
        alt: "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。"
      },
    ]
  },
  {
    title: "基礎工事各種",
    paragraphs: [
      {
        text: "鉄骨基礎",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "住宅基礎",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
    ]
  },
  {
    title: "改修工事",
    paragraphs: [
      {
        text: "塗装",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "足場",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "建具（内外装）",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "電気工事",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "管工事（エアコン等）",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
    ]
  },
  {
    title: "上下水道",
    paragraphs: [
      {
        text: "給水工事",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "下水工事",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },

    ]
  },
  {
    title: "公共工事",
    paragraphs: [
      {
        text: "道路工事",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "河川工事",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "水道工事",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "下水道工事",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
    ]
  },
  {
    title: "地域の困りごと解決",
    paragraphs: [
      {
        text: "真砂土敷き均し",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
      {
        text: "お祭り手伝い",
        imgSrc: placeholder,
        alt: "aaaaaaaaaaa",
      },
    ]
  },
]

export default function WorksPage() {
  return (
    <div className="font-genei">
      <div className="shadow-sides bg-white max-w-screen-lg mx-auto">
        <div className="w-auto h-[64px]"></div>
        <div className="max-w-screen-lg m-auto">
          <div className="block relative w-full mb-[var(--section-mb-sm)] md:mb-[var(--section-mb-md)]">
            <Image
              src={bang}
              alt="背景"
              sizes="100vw"
            />
            <Image
              src={works}
              alt="工事実績"
              width={250}
              objectPosition="center center"
              style={{margin: 'auto', position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%,-50%)'}}
            />
          </div>
          <div className="max-w-screen-md m-auto pr-8 pl-8">
            {items.map((c, i) => (
              <WorksItem key={i} {...c} />
            ))}
          </div>
        </div>
      </div>
      <MenuBar></MenuBar>
    </div>
  )
}