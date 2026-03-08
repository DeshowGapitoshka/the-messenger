import Sidebar from "../../components/Sidebar/Sidebar";
import styles from "./ChatPage.module.scss";

const ChatPage = () => {
  return (
    <main className={styles.chatPage}>
      <Sidebar />
      <div className={styles.body}></div>
    </main>
  );
};

export default ChatPage;
