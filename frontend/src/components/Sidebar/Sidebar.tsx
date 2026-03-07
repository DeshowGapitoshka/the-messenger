import styles from "./Sidebar.module.scss";
import SidebarProfile from "./SidebarProfile/SidebarProfile";
import SidebarButton from "./SidebarButton/SidebarButton";

import chats from "../../utils/icons/sidebar_chats.png";
import calls from "../../utils/icons/sidebar_calls.png";
import settings from "../../utils/icons/sidebar_settings.png";

const Sidebar = () => {
  return (
    <div className={styles.sidebar}>
      <div className={styles.top}>
        <SidebarProfile />
        <SidebarButton href="https://google.com" iconSrc={chats} />
        <SidebarButton href="https://google.com" iconSrc={calls} />
      </div>
      <div className={styles.bottom}>
        <SidebarButton href="https://google.com" iconSrc={settings} />
      </div>
    </div>
  );
};

export default Sidebar;
