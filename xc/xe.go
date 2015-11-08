package xc

/* 元素事件
   XE_ELEPROCE   1   元素处理过程事件.
   XE_PAINT   2   元素绘制事件
   XE_PAINT_END   3   该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END();
   XE_PAINT_SCROLLVIEW   4   滚动视图绘制事件.
   XE_MOUSEMOVE   5   元素鼠标移动事件.
   XE_MOUSESTAY   6   元素鼠标进入事件.
   XE_MOUSEHOVER   7   元素鼠标悬停事件.
   XE_MOUSELEAVE   8   元素鼠标离开事件.
   XE_MOUSEWHEEL   9   元素鼠标滚轮滚动事件.
   XE_LBUTTONDOWN   10   鼠标左键按下事件.
   XE_LBUTTONUP   11   鼠标左键弹起事件.
   XE_RBUTTONDOWN   12   鼠标右键按下事件.
   XE_RBUTTONUP   13   鼠标右键弹起事件.
   XE_LBUTTONDBCLICK   14   鼠标左键双击事件.
*/
const (
    XE_ELEPROCE = 1 + iota
    XE_PAINT
    XE_PAINT_END
    XE_PAINT_SCROLLVIEW
    XE_MOUSEMOVE
    XE_MOUSESTAY
    XE_MOUSEHOVER
    XE_MOUSELEAVE
    XE_MOUSEWHEEL
    XE_LBUTTONDOWN
    XE_LBUTTONUP
    XE_RBUTTONDOWN
    XE_RBUTTONUP
    XE_LBUTTONDBCLICK
)

/* 元素事件
   XE_SETFOCUS   31   元素获得焦点事件.
   XE_KILLFOCUS   32   元素失去焦点事件.
   XE_DESTROY   33   元素销毁事件.
   XE_BNCLICK   34   按钮点击事件.
   XE_BUTTON_CHECK   35   按钮选中事件.
   XE_SIZE   36   元素大小改变事件.
   XE_SHOW   37   元素显示隐藏事件.
   XE_SETFONT   38   元素设置字体事件.
   XE_KEYDOWN   39   元素按键事件.
   XE_CHAR   40   通过TranslateMessage函数翻译的字符事件.
*/
const (
    XE_SETFOCUS = 31 + iota
    XE_KILLFOCUS
    XE_DESTROY
    XE_BNCLICK
    XE_BUTTON_CHECK
    XE_SIZE
    XE_SHOW
    XE_SETFONT
    XE_KEYDOWN
    XE_CHAR
)

/* 元素事件
   XE_SETCAPTURE   51   元素设置鼠标捕获.
   XE_KILLCAPTURE   52   元素失去鼠标捕获.
   XE_SETCURSOR   53   设置鼠标光标
   XE_SCROLLVIEW_SCROLL_H   54   滚动视图元素水平滚动事件,滚动视图触发.
   XE_SCROLLVIEW_SCROLL_V   55   滚动视图元素垂直滚动事件,滚动视图触发.
   XE_SBAR_SCROLL   56   滚动条元素滚动事件,滚动条触发.
   XE_MENU_POPUP   57   菜单弹出
   XE_MENU_POPUP_WND   58   菜单弹出窗口
   XE_MENU_SELECT   59   弹出菜单项选择事件.
   XE_MENU_DRAW_BACKGROUND   60   绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
   XE_MENU_DRAWITEM   61   绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
   XE_MENU_EXIT   62   弹出菜单退出事件.
   XE_SLIDERBAR_CHANGE   63   滑动条元素,滑块位置改变事件.
   XE_PROGRESSBAR_CHANGE   64   进度条元素,进度改变事件.
*/
const (
    XE_SETCAPTURE = 51 + iota
    XE_KILLCAPTURE
    XE_SETCURSOR
    XE_SCROLLVIEW_SCROLL_H
    XE_SCROLLVIEW_SCROLL_V
    XE_SBAR_SCROLL
    XE_MENU_POPUP
    XE_MENU_POPUP_WND
    XE_MENU_SELECT
    XE_MENU_DRAW_BACKGROUND
    XE_MENU_DRAWITEM
    XE_MENU_EXIT
    XE_SLIDERBAR_CHANGE
    XE_PROGRESSBAR_CHANGE
)

/* 元素事件
   XE_COMBOBOX_SELECT   71   组合框下拉列表项选择事件.
   XE_COMBOBOX_POPUP_LIST   72   组合框下拉列表弹出事件.
   XE_COMBOBOX_EXIT_LIST   73   组合框下拉列表退出事件.
*/
const (
    XE_COMBOBOX_SELECT = 71 + iota
    XE_COMBOBOX_POPUP_LIST
    XE_COMBOBOX_EXIT_LIST
)

/* 元素事件
   XE_LISTBOX_TEMP_CREATE   81   列表框元素,项模板创建.
   XE_LISTBOX_TEMP_CREATE_END   82   列表框元素,项模板创建完成事件.
   XE_LISTBOX_TEMP_DESTROY   83   列表框元素,项模板销毁.
   XE_LISTBOX_TEMP_ADJUST_COORDINATE   84   列表框元素,项模板调整坐标.
   XE_LISTBOX_DRAWITEM   85   列表框元素,项绘制事件.
   XE_LISTBOX_SELECT   86   列表框元素,项选择事件.
*/
const (
    XE_LISTBOX_TEMP_CREATE = 81 + iota
    XE_LISTBOX_TEMP_CREATE_END
    XE_LISTBOX_TEMP_DESTROY
    XE_LISTBOX_TEMP_ADJUST_COORDINATE
    XE_LISTBOX_DRAWITEM
    XE_LISTBOX_SELECT
)

/* 元素事件
   XE_LIST_TEMP_CREATE   101   列表元素,项模板创建.
   XE_LIST_TEMP_CREATE_END   102   列表元素,项模板创建完成事件.
   XE_LIST_TEMP_DESTROY   103   列表元素,项模板销毁.
   XE_LIST_TEMP_ADJUST_COORDINATE   104   列表元素,项模板调整坐标.
   XE_LIST_DRAWITEM   105   列表元素,绘制项.
   XE_LIST_SELECT   106   列表元素,项选择事件.
   XE_LIST_HEADER_DRAWITEM   107   列表元素绘制列表头项.
*/
const (
    XE_LIST_TEMP_CREATE = 101 + iota
    XE_LIST_TEMP_CREATE_END
    XE_LIST_TEMP_DESTROY
    XE_LIST_TEMP_ADJUST_COORDINATE
    XE_LIST_DRAWITEM
    XE_LIST_SELECT
    XE_LIST_HEADER_DRAWITEM
)

/* 元素事件
   XE_TREE_TEMP_CREATE   121   树元素,项模板创建.
   XE_TREE_TEMP_CREATE_END   122   树元素,项模板创建完成.
   XE_TREE_TEMP_DESTROY   123   树元素,项模板销毁.
   XE_TREE_TEMP_ADJUST_COORDINATE   124   树元素,项模板,调整项坐标.
   XE_TREE_DRAWITEM   125   树元素,绘制项.
   XE_TREE_SELECT   126   树元素,项选择事件.
   XE_TREE_EXPAND   127   树元素,项展开收缩事件.
*/
const (
    XE_TREE_TEMP_CREATE = 121 + iota
    XE_TREE_TEMP_CREATE_END
    XE_TREE_TEMP_DESTROY
    XE_TREE_TEMP_ADJUST_COORDINATE
    XE_TREE_DRAWITEM
    XE_TREE_SELECT
    XE_TREE_EXPAND
)

/* 元素事件
   XE_LISTVIEW_TEMP_CREATE   141   列表视元素,项模板创建.
   XE_LISTVIEW_TEMP_CREATE_END   142   列表视元素,项模板创建完成事件.
   XE_LISTVIEW_TEMP_DESTROY   143   列表视元素,项模板销毁.
   XE_LISTVIEW_TEMP_ADJUST_COORDINATE   144   列表视元素,项模板调整坐标.
   XE_LISTVIEW_DRAWITEM   145   列表视元素,自绘项.
   XE_LISTVIEW_SELECT   146   列表视元素,项选择事件.
   XE_LISTVIEW_EXPAND   147   列表视元素,组展开收缩事件.
*/
const (
    XE_LISTVIEW_TEMP_CREATE = 141 + iota
    XE_LISTVIEW_TEMP_CREATE_END
    XE_LISTVIEW_TEMP_DESTROY
    XE_LISTVIEW_TEMP_ADJUST_COORDINATE
    XE_LISTVIEW_DRAWITEM
    XE_LISTVIEW_SELECT
    XE_LISTVIEW_EXPAND
)

/* 元素事件
   XE_PGRID_VALUE_CHANGE   161   属性网格元素 项值改变事件
   XE_RICHEDIT_CHANGE   162   富文本元素 用户修改内容事件,只有当用户操作时才会触发,需要开启后有效, XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE()； XRichEdit_SetText()、XRichEdit_InsertString()不会触发此事件
   XE_TABBAR_SELECT   221   TabBar标签按钮选择改变事件
   XE_TABBAR_DELETE   222   TabBar标签按钮删除事件
   XE_MONTHCAL_CHANGE   231   月历元素日期改变事件, 未开放
   XE_DATETIME_CHANGE   241   日期时间元素,内容改变事件, 未开放
   XE_DATETIME_POPUP_MONTHCAL   242   日期时间元素,弹出月历卡片事件, 未开放
   XE_DATETIME_EXIT_MONTHCAL   243   日期时间元素,弹出的月历卡片退出事件, 未开放
*/
const (
    XE_PGRID_VALUE_CHANGE      = 161
    XE_RICHEDIT_CHANGE         = 162
    XE_TABBAR_SELECT           = 221
    XE_TABBAR_DELETE           = 222
    XE_MONTHCAL_CHANGE         = 231
    XE_DATETIME_CHANGE         = 241
    XE_DATETIME_POPUP_MONTHCAL = 242
    XE_DATETIME_EXIT_MONTHCAL  = 243
)
