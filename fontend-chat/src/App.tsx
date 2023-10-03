import { Button, ButtonGroup, Card, Textarea } from "@nextui-org/react";
import { FiSend, FiPlus } from "react-icons/fi";

function App() {
  return (
    <main className="md:p-2 xl:p-5 w-screen h-screen">
      <Card className="mx-auto h-full  max-w-[768px]  sm:rounded-xl rounded-none  sm:p-2 md:p-5">
        <h1 className=" text-center">AirDrop-Go</h1>
        <div className="h-full flex flex-col">
          <div className="h-full">xxx</div>
          <div className="flex">
            <Textarea
              minRows={1}
              maxRows={1}
              labelPlacement="outside"
              placeholder="Enter your description"
              className="w-full"
            />
            <ButtonGroup variant="flat">
              <Button isIconOnly color="primary" aria-label="Like">
                <FiSend />
              </Button>
              <Button
                isIconOnly
                color="warning"
                variant="faded"
                aria-label="Take a photo"
              >
                <FiPlus />
              </Button>
              {/* <Dropdown placement="bottom-end">
                <DropdownTrigger>
                  <Button
                    isIconOnly
                    color="warning"
                    variant="faded"
                    aria-label="Take a photo"
                  >
                    <FiPlus />
                  </Button>
                </DropdownTrigger>
                <DropdownMenu
                  disallowEmptySelection
                  aria-label="Merge options"
                  selectionMode="single"
                  className="max-w-[300px]"
                >
                  <DropdownItem key="merge">x</DropdownItem>
                  <DropdownItem key="squash">xxx</DropdownItem>
                  <DropdownItem key="rebase">xxxxx</DropdownItem>
                </DropdownMenu>
              </Dropdown> */}
            </ButtonGroup>
          </div>
        </div>
      </Card>
    </main>
  );
}

export default App;
