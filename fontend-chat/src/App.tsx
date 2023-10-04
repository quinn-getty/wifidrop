import { Button, ButtonGroup, Card, Textarea } from "@nextui-org/react";
import { FiSend, FiPlus } from "react-icons/fi";
import {
  FaFileCsv,
  FaFileExcel,
  FaFileImage,
  FaFilePdf,
  FaFilePowerpoint,
  FaFileVideo,
  FaFileWord,
  FaFileZipper,
  FaFile,
} from "react-icons/fa6";
import {
  BiSolidFileCss,
  BiSolidFileHtml,
  BiSolidFileJs,
  BiSolidFileJson,
  BiSolidFileTxt,
} from "react-icons/bi";

import { useEffect, useRef, useState } from "react";
import { http } from "./utils/http";
import dayjs from "dayjs";
import { getWsClient } from "./utils/ws_client";

function getRandomColor1() {
  // 生成随机红色通道值
  const red = getRandomValue(50, 200);
  // 生成随机绿色通道值
  const green = getRandomValue(40, 210);
  // 生成随机蓝色通道值
  const blue = getRandomValue(0, 100);

  // 将通道值组合成颜色字符串
  const color = `rgb(${red}, ${green}, ${blue})`;

  return color;
}

function getRandomValue(min: number, max: number) {
  // 生成介于min和max之间的随机整数
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function App() {
  const inputRef = useRef<HTMLInputElement>(null);
  const sendMessage = () => {
    if (!inputRef.current?.value) {
      return;
    }
    const params = {
      content: inputRef.current?.value,
      type: "text",
    };
    console.log(params);
    http.post("/api/v2/send", params);
  };

  const fileInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target?.files?.[0];
    if (!file) return;
    const formData = new FormData();
    formData.append("raw", file);
    console.log(file.type);

    formData.append("type", file.type || "file");
    http({
      method: "post",
      url: "/api/v2/upload",
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  };

  const getHistory = async () => {
    const res = await http.get("/api/v2/history");
    setList(res.data.list || []);
  };

  const [list, setList] = useState<
    {
      type: string;
      content: string;
      time: number;
      ip: string;
    }[]
  >([]);

  useEffect(() => {
    getHistory();
  }, []);

  useEffect(() => {
    getWsClient().then((c) => {
      c.onMessage((data) => {
        console.log(data);
      });
    });
  }, []);

  return (
    <main className="md:p-2 xl:p-5 w-screen h-screen">
      <Card className="mx-auto h-full  max-w-[768px]  sm:rounded-xl rounded-none  sm:p-2 md:p-5">
        <h1 className=" text-center">AirDrop-Go</h1>
        <div className="h-full flex flex-col">
          <div className="h-full overflow-auto w-auto pb-4">
            {list.map((item, index) => (
              <div key={index} className="w-full mb-4 flex flex-wrap">
                <ItemCard item={item} />
              </div>
            ))}
          </div>
          <div className="flex pb-2">
            <Textarea
              ref={inputRef}
              minRows={1}
              maxRows={1}
              labelPlacement="outside"
              placeholder="Enter your description"
              className="w-full"
            />
            <ButtonGroup variant="flat">
              <Button
                isIconOnly
                color="primary"
                aria-label="Like"
                onClick={sendMessage}
              >
                <FiSend />
              </Button>
              <Button
                as={"label"}
                isIconOnly
                color="warning"
                variant="faded"
                aria-label="Take a photo"
              >
                <FiPlus />
                <input type="file" hidden onChange={fileInputChange} />
              </Button>
            </ButtonGroup>
          </div>
        </div>
      </Card>
    </main>
  );
}

const icons = {
  pdf: <FaFilePdf />,
  csv: <FaFileCsv />,
  excel: <FaFileExcel />,
  img: <FaFileImage />,
  ppt: <FaFilePowerpoint />,
  video: <FaFileVideo />,
  word: <FaFileWord />,
  zip: <FaFileZipper />,
  file: <FaFile />,
};

icons;

const getIcon = (type: string) => {
  if (type === "application/pdf") {
    return <FaFilePdf />;
  }
  if (type === "application/zip") {
    return <FaFileZipper />;
  }
  if (type === "application/vnd.ms-powerpoint") {
    return <FaFilePowerpoint />;
  }
  if (type === "text/csv") {
    return <FaFileCsv />;
  }
  if (type === "text/html") {
    return <BiSolidFileHtml />;
  }
  if (type === "text/javascript") {
    return <BiSolidFileJs />;
  }
  if (type === "application/json") {
    return <BiSolidFileJson />;
  }
  if (type === "text/css") {
    return <BiSolidFileCss />;
  }
  if (type === "text/plain") {
    return <BiSolidFileTxt />;
  }
  if (/^video\//.test(type)) {
    return <FaFileVideo />;
  }
  if (/excel$/.test(type)) {
    return <FaFileExcel />;
  }
  if (/word/.test(type)) {
    return <FaFileWord />;
  }

  return <FaFile />;
};

const ItemCard = ({
  item,
}: {
  item: {
    type: string;
    content: string;
    time: number;
    ip: string;
  };
}) => {
  return (
    <div className="max-w-[90%]">
      <div className="text-sm">
        <span className="text-[#333333] font-[600]">{item.ip}</span>
        <span className="text-xs pl-2 text-[#d1d1d1]">
          {dayjs(item.time).format("YYYY-MM-DD HH:mm:ss")}
        </span>
      </div>
      <Card className="shadow-sm my-1 bg-[#dfdfdf] px-4 py-4">
        {item.type !== "string" ? (
          <>
            <div className="flex">
              {/^image\//.test(item.type) ? (
                <img
                  className="max-w-[80%] max-h-[250px] w-auto"
                  src={`/api/v2/download/${item.content}`}
                />
              ) : (
                <>
                  <div
                    className={`text-[80px]`}
                    style={{
                      color: getRandomColor1(),
                    }}
                  >
                    {getIcon(item.type)}
                  </div>
                  <div className="pl-2">
                    <a href={`/api/v2/download/${item.content}`}>
                      <h1>{item.content.replace(/^([^_]*)_/, "")}</h1>
                    </a>
                  </div>
                </>
              )}
            </div>
          </>
        ) : (
          <>{item.content}</>
        )}
      </Card>
    </div>
  );
};

export default App;
