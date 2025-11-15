'use client';

import Image from 'next/image';
import MenuBar from "@/app/components/MenuBar";

// import { noto, genei } from "./"

import aboutUs from '../../public/text_about_us.webp';
import bang from '../../public/bang.webp';
import header from '../../public/header.webp';
import tecchan from '../../public/tecchan.webp'

export default function AboutPage() {
  return (
    <div className="font-genei">
      <div className="shadow-sides bg-white max-w-screen-lg mx-auto">
        <div className="w-auto h-[64px]"></div>
        <div className="max-w-screen-lg m-auto">
          <div className="block relative w-full">
            <Image
              src={bang}
              alt="背景"
              sizes="100vw"
            />
            <Image
              src={aboutUs}
              alt="弊社について"
              width={300}
              objectPosition="center center"
              style={{margin: 'auto', position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%,-50%)'}}
            />
          </div>
          <div className="max-w-screen-md m-auto pr-8 pl-8">
            <div className="mt-8 mb-4">
              <Image
                src={header}
                alt="ヘッダ"
                width={32}
                height={32}
                style={{display: "inline", marginRight: '10px'}}
              />
              <span className="align-[-6px] text-2xl font-bold">企業情報</span>
            </div>

            <div className="w-full relative rounded-xl border border-gray-400">
              <table className="w-full text-sm text-left rtl:text-right text-body">
                <tbody>

                <tr className="border-b border-gray-400 border-default">
                  <td className="px-6 py-3">設立年</td>
                  <td className="px-6 py-3">1982年</td>
                </tr>

                <tr className="border-b border-gray-400 border-default">
                  <td className="px-6 py-3">従業員数</td>
                  <td className="px-6 py-3">5人</td>
                </tr>

                <tr className="border-b border-gray-400 border-default">
                  <td className="px-6 py-3">資本金</td>
                  <td className="px-6 py-3">3,000万円</td>
                </tr>

                <tr className="border-b border-gray-400 border-default">
                  <td className="px-6 py-3">事業内容</td>
                  <td className="px-6 py-3">土木工事、建築工事、上下水道工事 全般</td>
                </tr>

                <tr className="border-b border-gray-400 border-default">
                  <td className="px-6 py-3">法人番号</td>
                  <td className="px-6 py-3">626000201498236</td>
                </tr>

                <tr>
                  <td className="px-6 py-3">代表者名</td>
                  <td className="px-6 py-3">中原 哲哉</td>
                </tr>

                </tbody>
              </table>
            </div>

            <div className="mt-16 mb-4">
              <Image
                src={header}
                alt="ヘッダ"
                width={32}
                height={32}
                style={{display: "inline", marginRight: '10px'}}
              />
              <span className="align-[-6px] text-2xl font-bold">沿革</span>
            </div>
            <ul className="list-disc ml-[1.5em]">
              <li>1982年6月 ─ 有限会社中原建設工業 創業</li>
              <li>1922年6月 ─ 創業40年達成</li>
            </ul>
            <div className="mt-16 mb-4">
              <div className="mb-8">
                <Image
                  src={header}
                  alt="ヘッダ"
                  width={32}
                  height={32}
                  style={{display: "inline", marginRight: '10px'}}
                />
                <span className="align-[-6px] text-2xl font-bold">社長より</span>
              </div>
              <div className="mb-8 flex">
                <Image src={tecchan} alt="社長" width={150} height={200} className="mr-10"/>
                <div className="
                  relative flex flex-grow-[1] items-center justify-center h-[4rem] m-auto box-border mr-[8px] z-0

                  before:content-[''] before:absolute before:inset-0
                  before:bg-white before:border before:border-black
                  before:z-0

                  after:content-[''] after:absolute after:-bottom-[8px] after:-right-[8px]
                  after:w-full after:h-full
                  after:bg-[radial-gradient(#000_1px,transparent_1px),radial-gradient(#000_1px,transparent_1px)]
                  after:bg-[length:6px_6px]
                  after:[background-position:-1px_-1px,2px_2px]
                  after:-z-10
                ">
                  <span className="relative font-bold">こんにちは！社長の中原哲哉です！</span>
                </div>
              </div>
              <p className="mb-4">
                おかげさまで、地元赤磐市で40年営業を続けています。
                岡山県東備地域の土木工事、赤磐市の土木・建築・水道・下水工事、民間施設の建築・土木修繕工事、一般住宅の新築・リフォーム工事、草刈り業務（ラジコンによる）など、いろいろな分野の作業を受注しています。
              </p>
              <p className="mb-4">
                仕事の目的は、人生を楽しむこと！そしてその為に、みんなで楽しく仕事をこなしてしっかり利益を出していくことをモットーにしています。
              </p>
              <p className="mb-4">
                今日を楽しめる方、連絡待ってます。是非一緒に働きましょう！
              </p>
            </div>
          </div>
        </div>
      </div>
      <MenuBar></MenuBar>
    </div>
  );
}
