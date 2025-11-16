'use client';

import Image from 'next/image';
import MenuBar from "@/app/components/MenuBar";

import inquiry from '../../public/text_inquiry.webp';
import bang from '../../public/bang.webp';
import header from '../../public/header.webp';

export default function RecruitPage() {
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
              src={inquiry}
              alt="お問い合わせ"
              width={300}
              objectPosition="center center"
              style={{margin: 'auto', position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%,-50%)'}}
            />
          </div>
          <div className="max-w-screen-md m-auto pr-8 pl-8">
            <div className="mb-[var(--section-mb-sm)] md:mb-[var(--section-mb-md)]">
              <div className="mb-4">
                <Image
                  src={header}
                  alt="ヘッダ"
                  width={32}
                  height={32}
                  style={{display: "inline", marginRight: '10px'}}
                />
                <span className="align-[-6px] text-2xl font-bold">連絡先</span>
              </div>

              <div className="mb-4">
                <p className="text-sm">お気軽にご連絡ください。</p>
              </div>

              <div className="w-full relative mb-8 rounded-xl border border-gray-400">
                <table className="w-full text-sm text-left rtl:text-right text-body">
                  <colgroup>
                    <col className="w-1/4"/>
                  </colgroup>
                  <tbody>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">電話番号</td>
                    <td className="px-6 py-3 underline decoration-dotted"><a href="tel:0869541700">086-954-1700</a></td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">FAX</td>
                    <td className="px-6 py-3">086-954-1702</td>
                  </tr>
                  <tr>
                    <td className="px-6 py-3">メール</td>
                    <td className="px-6 py-3 underline decoration-dotted"><a href="mailto:nakahara@nakaken.0am.jp">nakahara &lt;at&gt; nakaken.0am.jp</a></td>
                  </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
      <MenuBar></MenuBar>
    </div>
  );
}
