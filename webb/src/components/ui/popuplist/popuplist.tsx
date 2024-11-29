import { ReactNode } from "react";

interface listItem {
    key: string;
    icon: ReactNode;
}

interface popupList {
    width: number;
    top: listItem[];
    bottom: Record<string, ReactNode>[];
}

export function PopupList({top, bottom}: popupList): JSX.Element {
    return (
        <div>
            <div>
                {
                    top.map((value, i) => {
                        return (
                            <div key={i}>
                                <div>{value.icon}</div>
                                <div>{value.key}</div>
                            </div>
                        )
                    })
                }
            </div>

            <hr/>

            <div>
                {
                    bottom.map((value, i) => {
                        return (
                            <div key={i}>
                                <div>{value.icon}</div>
                                <div>{value.key}</div>
                            </div>
                        )
                    })
                }
            </div>
        </div>
    )
}