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
import { useRef } from "react";
import { http } from "./utils/http";

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
    http({
      method: "post",
      url: "/api/v2/upload",
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  };
  return (
    <main className="md:p-2 xl:p-5 w-screen h-screen">
      <Card className="mx-auto h-full  max-w-[768px]  sm:rounded-xl rounded-none  sm:p-2 md:p-5">
        <h1 className=" text-center">AirDrop-Go</h1>
        <div className="h-full flex flex-col">
          <div className="h-full overflow-auto w-auto pb-4">
            {[
              1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2,
              3, 4, 5, 6, 7, 8, 9, 0,
            ].map((_, index) => (
              <div key={index} className="w-full flex flex-wrap">
                <div className="max-w-[90%]">
                  <div className="text-sm">
                    <span className="text-[#333333] font-[600]">
                      192.17.1.1
                    </span>
                    <span className="text-xs pl-2 text-[#d1d1d1]">
                      2020.2.2
                    </span>
                  </div>
                  <Card className="shadow-sm my-1 bg-[#dfdfdf] px-2 py-2.5">
                    <div className="flex">
                      <div
                        className={`text-[40px]`}
                        style={{
                          color: getRandomColor1(),
                        }}
                      >
                        <FaFilePdf />
                      </div>
                      <div className="pl-2">
                        <h1>这是一个pdf</h1>
                      </div>
                    </div>
                  </Card>
                </div>
              </div>
            ))}
          </div>
          <div className="flex">
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

export default App;
