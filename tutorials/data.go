package tutorials

import (
	page2 "bigbrother/tutorials/page"
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

type TutorialNode struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
	children     []TutorialNode
}

var (
	TutorialNodeList = []TutorialNode{
		{"Welcome", "", welcomeScreen, true, nil},
		{"MyTool", "", welcomeScreen, true, []TutorialNode{

			{Title: "Jwt解析", Intro: "jwt解析", View: page2.JwtParsePage},
			{Title: "UrlEncode", Intro: "UrlEncode", View: page2.UrlEncodePage},
			{Title: "时间转化", Intro: "TimePage", View: page2.TimePage},
			{Title: "Json2yaml", Intro: "Json2yaml", View: page2.Json2yaml},
			{Title: "unicode转中文", Intro: "unicode转中文", View: page2.UnicodePage},
			{Title: "驼峰转下滑线", Intro: "Camel2snake", View: page2.Camel2snake},
			{Title: "Base64Decode", Intro: "Base64Page", View: page2.Base64Page},
			{Title: "公约数", Intro: "公约数", View: page2.GcdPage},
			{Title: "ClockPage", Intro: "ClockPage", View: page2.ClockPage},
			{Title: "button", Intro: "time", View: page2.GetTimeButton},
			{Title: "About", Intro: "About", View: page2.About},
		}},
		{"Canvas",
			"See the canvas capabilities.",
			canvasScreen,
			true,
			nil,
		}, {"Animations",
			"See how to animate components.",
			makeAnimationScreen,
			true,
			nil,
		},
		{"Theme Icons",
			"Browse the embedded icons.",
			iconScreen,
			true,
			nil,
		},
		{"Widgets",
			"In this section you can see the features available in the toolkit widget set.\n" +
				"Expand the tree on the left to browse the individual tutorial elements.",
			widgetScreen,
			true,
			[]TutorialNode{
				{"Accordion",
					"Expand or collapse content panels.",
					makeAccordionTab,
					true,
					nil,
				},
				{"Button",
					"Simple widget for user tap handling.",
					makeButtonTab,
					true,
					nil,
				},
				{"Card",
					"Group content and widgets.",
					makeCardTab,
					true,
					nil,
				},
				{"Entry",
					"Different ways to use the entry widget.",
					makeEntryTab,
					true,
					nil,
				},
				{"Form",
					"Gathering input widgets for data submission.",
					makeFormTab,
					true,
					nil,
				},
				{"Input",
					"A collection of widgets for user input.",
					makeInputTab,
					true,
					nil,
				},
				{"Progress",
					"Show duration or the need to wait for a task.",
					makeProgressTab,
					true,
					nil,
				},
				{"Text",
					"Text handling widgets.",
					makeTextTab,
					true,
					nil,
				},
				{"Toolbar",
					"A row of shortcut icons for common tasks.",
					makeToolbarTab,
					true,
					nil,
				},
			},
		},
		{"Collections",
			"Collection widgets provide an efficient way to present lots of content.\n" +
				"The List, Table, and Tree provide a cache and re-use mechanism that make it possible to scroll through thousands of elements.\n" +
				"Use this for large data sets or for collections that can expand as users scroll.",
			collectionScreen,
			true,
			[]TutorialNode{
				{"List",
					"A vertical arrangement of cached elements with the same styling.",
					makeListTab,
					true,
					nil,
				},
				{"Table",
					"A two dimensional cached collection of cells.",
					makeTableTab,
					true,
					nil,
				},
				{"Tree",
					"A tree based arrangement of cached elements with the same styling.",
					makeTreeTab,
					true,
					nil,
				}},
		},
		{"Containers",
			"Containers group other widgets and canvas objects, organising according to their layout.\n" +
				"Standard containers are illustrated in this section, but developers can also provide custom " +
				"layouts using the fyne.NewContainerWithLayout() constructor.",
			containerScreen,
			true,
			[]TutorialNode{{"AppTabs",
				"A container to help divide up an application into functional areas.",
				makeAppTabsTab,
				true,
				nil,
			},
				{"Border",
					"A container that positions items around a central content.",
					makeBorderLayout,
					true,
					nil,
				},
				{"Box",
					"A container arranges items in horizontal or vertical list.",
					makeBoxLayout,
					true,
					nil,
				},
				{"Center",
					"A container to that centers child elements.",
					makeCenterLayout,
					true,
					nil,
				},
				{"DocTabs",
					"A container to display a single document from a set of many.",
					makeDocTabsTab,
					true,
					nil,
				},
				{"Grid",
					"A container that arranges all items in a grid.",
					makeGridLayout,
					true,
					nil,
				},
				{"Split",
					"A split container divides the container in two pieces that the user can resize.",
					makeSplitTab,
					true,
					nil,
				},
				{"Scroll",
					"A container that provides scrolling for it's content.",
					makeScrollTab,
					true,
					nil,
				}},
		},
		{"Dialogs",
			"Work with dialogs.",
			dialogScreen,
			true,
			nil,
		},
		{"Windows",
			"Window function demo.",
			windowScreen,
			false,
			nil,
		},
		{"Data Binding",
			"Connecting widgets to a data source.",
			bindingScreen,
			true,
			nil,
		},
		{"Advanced",
			"Debug and advanced information.",
			advancedScreen,
			true,
			nil,
		},
	}

	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{}
)

func convertNodesToTutorials(nodes []TutorialNode, parentKey string, tutorials map[string]Tutorial, nodeMap map[string][]string) {

	for _, node := range nodes {
		tutorial := Tutorial{
			Title:      node.Title,
			Intro:      node.Intro,
			View:       node.View,
			SupportWeb: node.SupportWeb,
		}
		tutorials[node.Title] = tutorial
		nodeMap[parentKey] = append(nodeMap[parentKey], node.Title)
		if node.children != nil && len(node.children) != 0 {
			convertNodesToTutorials(node.children, node.Title, tutorials, nodeMap)
		}
	}

}

func init() {
	convertNodesToTutorials(TutorialNodeList, "", Tutorials, TutorialIndex)
}
