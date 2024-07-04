document.addEventListener("DOMContentLoaded", () => {
    document.querySelectorAll(".主體-格局-內容-單個範例-文字-標題-切換原始碼").forEach(v => {
        v.addEventListener("click", () => {
            v.closest(".主體-格局-內容-單個範例").classList.toggle("主體-格局-內容-單個範例_檢視原始碼中")
        })
    })

    document.querySelectorAll(".頁腳-導航列-項目_回到頂部, .主體-格局-索引-回到頂部").forEach(v => {
        v.addEventListener("click", () => {
            document.body.scrollTop = 0
            document.documentElement.scrollTop = 0
            // 移除 Hash。
            history.pushState("", document.title, window.location.pathname + window.location.search)
        })
    })

    document.addEventListener("click", e => {
        document.querySelectorAll(".主體-格局-內容-工具列-項目-下拉式選單").forEach(e => {
            e.querySelector(".主體-格局-內容-工具列-項目-下拉式選單-群組").classList.remove("主體-格局-內容-工具列-項目-下拉式選單-群組_開啟的")
        })
    })

    document.querySelector(".穹頂-導航列-項目_選單按鈕")?.addEventListener("click", () => {
        document.querySelector(".主體-格局-導覽").classList.add("主體-格局-導覽_啟用的")
        document.querySelector(".主體-格局-遮罩").classList.add("主體-格局-遮罩_啟用的")
    })

    document.querySelector(".主體-格局-導覽-關閉按鈕")?.addEventListener("click", () => {
        document.querySelector(".主體-格局-導覽").classList.remove("主體-格局-導覽_啟用的")
        document.querySelector(".主體-格局-遮罩").classList.remove("主體-格局-遮罩_啟用的")
    })

    document.querySelectorAll(".主體-格局-內容-工具列-項目-下拉式選單").forEach(e => {
        e.addEventListener("click", e => {
            var menu = e.target.closest(".主體-格局-內容-工具列-項目-下拉式選單")

            if (e.target.classList.contains("主體-格局-內容-工具列-項目-下拉式選單-群組-項目")) {
                let [removes, add] = e.target.getAttribute("data-value").split(";")
                document.querySelector("html").classList.remove(...removes.split(","))
                if (add !== undefined) {
                    document.querySelector("html").classList.add(add)
                }
                menu.querySelectorAll(".主體-格局-內容-工具列-項目-下拉式選單-群組-項目").forEach(e => {
                    e.classList.remove("主體-格局-內容-工具列-項目-下拉式選單-群組-項目_啟用的")
                })
                menu.querySelector(".主體-格局-內容-工具列-項目-下拉式選單-文字-標籤").innerText = e.target.innerText
                e.target.classList.add("主體-格局-內容-工具列-項目-下拉式選單-群組-項目_啟用的")
            } else {
                setTimeout(() => {
                    menu.querySelector(".主體-格局-內容-工具列-項目-下拉式選單-群組").classList.add("主體-格局-內容-工具列-項目-下拉式選單-群組_開啟的")
                }, 1)
            }
        })
    })

    // document.querySelectorAll("[data-value]").forEach(e => {
    //     e.addEventListener("click", e => {
    //         let [removes, add] = e.target.getAttribute("data-value").split(";")
    //         document.querySelector("html").classList.remove(...removes.split(","))
    //         if (add !== undefined) {
    //             document.querySelector("html").classList.add(add);
    //         }
    //     });
    // });

    window.addEventListener("hashchange", rescanAnchor)
    rescanAnchor()
})

function rescanAnchor() {
    let u = document.location.hash
    if (u.length === 0) {
        return
    }
    document.querySelectorAll(".主體-格局-內容-單個範例-文字-標題").forEach(v => {
        v.classList.remove("主體-格局-內容-單個範例-文字-標題_被提及的")
    })

    document
        .querySelector(`a[id='${decodeURI(u.substring(u.indexOf("#") + 1))}']`)
        ?.closest(".主體-格局-內容-單個範例")
        .querySelector(".主體-格局-內容-單個範例-文字-標題")
        .classList.add("主體-格局-內容-單個範例-文字-標題_被提及的")

    document.querySelector(`a[id='${decodeURI(u.substring(u.indexOf("#") + 1))}']`)?.scrollIntoView()
}

function download(event) {
    const link = document.createElement("a")
    const file = new Blob([event.target.closest(".主體-格局-內容-單個範例-附帶程式碼_可下載的").querySelector("pre").innerText], { type: "text/plain" })
    link.href = URL.createObjectURL(file)
    link.download = "source-code.txt"
    link.click()
    URL.revokeObjectURL(link.href)
}
