import { type FC } from "react";
import styles from "./SidebarButton.module.scss";

interface ISidebarButton {
  href: string;
  iconSrc: string;
}

// CHANGE | alt
const SidebarButton: FC<ISidebarButton> = ({ href, iconSrc }) => {
  return (
    <a href={href} className={styles.btn}>
      <img src={iconSrc} className={styles.img} alt="The sidebar icon" />
    </a>
  );
};

export default SidebarButton;
