package info.sqldev.user;

public class UserPage {
    int count;
    int page_no;
    int page_size;
    User[] list;

    public int getCount() {
        return count;
    }

    public void setCount(int count) {
        this.count = count;
    }

    public int getPage_no() {
        return page_no;
    }

    public void setPage_no(int page_no) {
        this.page_no = page_no;
    }

    public int getPage_size() {
        return page_size;
    }

    public void setPage_size(int page_size) {
        this.page_size = page_size;
    }

    public User[] getList() {
        return list;
    }

    public void setList(User[] list) {
        this.list = list;
    }
}
