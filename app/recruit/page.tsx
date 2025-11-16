'use client';

import Image from 'next/image';
import MenuBar from "@/app/components/MenuBar";

import recruit from '../../public/text_recruit.webp';
import bang from '../../public/bang.webp';
import header from '../../public/header.webp';
import entry from '../../public/entry.webp';

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
              src={recruit}
              alt="採用情報"
              width={225}
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
                <span className="align-[-6px] text-2xl font-bold">仕事内容</span>
              </div>
              <div className="w-full relative rounded-xl border border-gray-400">
                <table className="w-full text-sm text-left rtl:text-right text-body">
                  <colgroup>
                    <col className="w-1/4"/>
                  </colgroup>

                  <tbody>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">職種</td>
                    <td className="px-6 py-3">建設業 土木・建築 監理及び作業</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">仕事内容</td>
                    <td className="px-6 py-3 text-justify">
                      <p className="mb-[1em]">
                        主な業務内容は、赤磐市内の土木工事・建築工事の現場管理技術者および現場技能者に向けた
                        現場書類・製図、見積りの作成です。近所の困りごとへの協力や、近所のお祭り準備作業等も業務に含まれます。
                        実際の内容はご本人に合わせて採用以降決めていきます。
                      </p>

                      <p className="mb-[1em]">
                        経験よりやる気のある方を募集しています。経験者も未経験者もやる気があれば大歓迎です。
                        PC 作業から重機のオペレータまで、出来るようにレクチャーいたします。
                        必要に応じて免許等は就業時間内に取得して頂きます。
                      </p>

                      <p className="mb-[1em]">
                        夏は暑くて、冬は寒い環境ですが、みんなで楽しく出来る会社を目指しています。
                      </p>

                      <p>
                        今日を楽しめる方、連絡をお待ちしております。是非一緒に働きましょう！<br/>
                      </p>
                    </td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">雇用形態</td>
                    <td className="px-6 py-3">正社員</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">雇用期間</td>
                    <td className="px-6 py-3">定めなし</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">就業場所</td>
                    <td className="px-6 py-3">事業所所在地と同じ（現場は主に赤磐市・和気町・備前市）
                    </td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">受動喫煙対策</td>
                    <td className="px-6 py-3">あり（屋内禁煙）電子たばこのみ室内喫煙可</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">マイカー通勤</td>
                    <td className="px-6 py-3">可、駐車場あり</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">年齢</td>
                    <td className="px-6 py-3">69歳以下、定年年齢を上限
                    </td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">学歴</td>
                    <td className="px-6 py-3">不問</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">経験</td>
                    <td className="px-6 py-3">土木管理・建築管理・土木技能者（いずれもあれば尚可）</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">PCスキル</td>
                    <td className="px-6 py-3">ワ－ド・エクセル（あれば尚可）<br/>建設ＣＡＤが使えるならもっと可
                    </td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">免許・資格</td>
                    <td className="px-6 py-3">普通自動車運転免許 <span
                      className="font-bold">必須（AT限定不可）</span><br/>
                      ２級土木施工管理技士 あれば尚可<br/>
                      ２級建築施工管理技士 あれば尚可<br/>
                      その他の土木・舗装・線路工事関係資格 あれば尚可<br/>
                    </td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">その他優遇等</td>
                    <td className="px-6 py-3">
                      当社はうらじゃを応援しているので、うらじゃの踊り子さんを優遇いたします。<br/><br/>
                      1. うらじゃ衣装代は当社負担（勤続６ヶ月以上から）<br/>
                      2. 繁忙期でも本祭、遠征を優先してOK<br/>
                    </td>
                  </tr>
                  <tr>
                    <td className="px-6 py-3">試用期間</td>
                    <td className="px-6 py-3">3 ヶ月、期間中は手当無し</td>
                  </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <div className="mb-[var(--section-mb-sm)] md:mb-[var(--section-mb-md)]">
              <div className="mb-4">
                <Image
                  src={header}
                  alt="ヘッダ"
                  width={32}
                  height={32}
                  style={{display: "inline", marginRight: '10px'}}
                />
                <span className="align-[-6px] text-2xl font-bold">賃金・手当</span>
              </div>
              <div className="w-full relative rounded-xl border border-gray-400">
                <table className="w-full text-sm text-left rtl:text-right text-body">
                  <colgroup>
                    <col className="w-1/4"/>
                  </colgroup>
                  <tbody>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">月額</td>
                    <td className="px-6 py-3">192,000 円 〜 312,000 円<br/>（月平均労働日数23.3日）</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">賃金形態</td>
                    <td className="px-6 py-3">日給 8,000 円 〜 13,000 円</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">通勤手当</td>
                    <td className="px-6 py-3">実費支給（上限なし）</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">賃金締切日</td>
                    <td className="px-6 py-3">月末</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">賃金支払日</td>
                    <td className="px-6 py-3">翌月 5 日</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">昇給</td>
                    <td className="px-6 py-3">あり<br/>1 月あたり 10,000 円 〜 20,000 円</td>
                  </tr>
                  <tr>
                    <td className="px-6 py-3">賞与</td>
                    <td className="px-6 py-3">あり 年2回<br/>賞与金額 300,000 円 〜 500,000 円</td>
                  </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <div className="mb-[var(--section-mb-sm)] md:mb-[var(--section-mb-md)]">
              <div className="mb-4">
                <Image
                  src={header}
                  alt="ヘッダ"
                  width={32}
                  height={32}
                  style={{display: "inline", marginRight: '10px'}}
                />
                <span className="align-[-6px] text-2xl font-bold">労働時間</span>
              </div>
              <div className="w-full relative rounded-xl border border-gray-400">
                <table className="w-full text-sm text-left rtl:text-right text-body">
                  <colgroup>
                    <col className="w-1/4"/>
                  </colgroup>
                  <tbody>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">就業時間</td>
                    <td className="px-6 py-3">8時00分 〜 17時00分 7時間程度</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">休憩</td>
                    <td className="px-6 py-3">10:00〜10:30<br/>12:00〜13:00<br/>15:00〜15:30</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">時間外労働</td>
                    <td className="px-6 py-3">なし</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">年間休日数</td>
                    <td className="px-6 py-3">85日</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">休日</td>
                    <td className="px-6 py-3">第2・第4土曜日<br/>年末年始<br/>盆休み<br/>その他会社が指定する日</td>
                  </tr>
                  <tr>
                    <td className="px-6 py-3">6ヶ月経過後の<br/>年次有給休暇日数</td>
                    <td className="px-6 py-3">10 日</td>
                  </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <div className="mb-[var(--section-mb-sm)] md:mb-[var(--section-mb-md)]">
              <div className="mb-4">
                <Image
                  src={header}
                  alt="ヘッダ"
                  width={32}
                  height={32}
                  style={{display: "inline", marginRight: '10px'}}
                />
                <span className="align-[-6px] text-2xl font-bold">その他の労働条件等</span>
              </div>
              <div className="w-full relative rounded-xl border border-gray-400">
                <table className="w-full text-sm text-left rtl:text-right text-body">
                  <colgroup>
                    <col className="w-1/4"/>
                  </colgroup>
                  <tbody>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">加入保険</td>
                    <td className="px-6 py-3">雇用、労災、健康、厚生</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">定年</td>
                    <td className="px-6 py-3">70 歳</td>
                  </tr>
                  <tr>
                    <td className="px-6 py-3">再雇用制度</td>
                    <td className="px-6 py-3">あり（上限 75 歳まで）</td>
                  </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <div className="mb-[var(--section-mb-sm)] md:mb-[var(--section-mb-md)]">
              <div className="mb-4">
                <Image
                  src={header}
                  alt="ヘッダ"
                  width={32}
                  height={32}
                  style={{display: "inline", marginRight: '10px'}}
                />
                <span className="align-[-6px] text-2xl font-bold">選考</span>
              </div>
              <div className="w-full relative mb-8 rounded-xl border border-gray-400">
                <table className="w-full text-sm text-left rtl:text-right text-body">
                  <colgroup>
                    <col className="w-1/4"/>
                  </colgroup>
                  <tbody>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">選考方法</td>
                    <td className="px-6 py-3">面接（1 回）</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">選考結果通知</td>
                    <td className="px-6 py-3">7 日以内に電話で通知</td>
                  </tr>
                  <tr>
                    <td className="px-6 py-3">応募書類</td>
                    <td className="px-6 py-3">履歴書（面接時持参）</td>
                  </tr>
                  </tbody>
                </table>
              </div>
            </div>

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
              <div className="w-full relative mb-8 rounded-xl border border-gray-400">
                <table className="w-full text-sm text-left rtl:text-right text-body">
                  <colgroup>
                    <col className="w-1/4"/>
                  </colgroup>
                  <tbody>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">担当者</td>
                    <td className="px-6 py-3">中原 哲哉</td>
                  </tr>
                  <tr className="border-b border-gray-400 border-default">
                    <td className="px-6 py-3">電話番号</td>
                    <td className="px-6 py-3 underline decoration-dotted"><a href="tel:0869541700">086-954-1700</a></td>
                  </tr>
                  <tr>
                    <td className="px-6 py-3">FAX</td>
                    <td className="px-6 py-3">086-954-1702</td>
                  </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <div className="mb-16">
              <p className="text-sm">※上記はハローワークに公開した 2025 年 9 月 3
                日現在の求人票の内容を元にしています。</p>
            </div>

            <a href="tel:0869541700">
              <div className="relative mb-8 hover:scale-105 transition duration-100 cursor-pointer">
                <Image src={entry} alt="エントリーはこちら"/>
                <div className="absolute inset-0 mb-4 pt-2 flex items-center justify-center">
                  <span className="text-xl sm:text-4xl font-bold text-center">
                    採用担当者に<br/>電話する
                    <span className="inline-block  rotate-[15deg] translate-x-[0.05em] translate-y-[0em]">!</span>
                    <span className="inline-block  rotate-[15deg] translate-x-[0.05em] translate-y-[0em]">!</span>
                  </span>
                </div>
              </div>
            </a>
          </div>
        </div>
      </div>
      <MenuBar></MenuBar>
    </div>
  );
}
