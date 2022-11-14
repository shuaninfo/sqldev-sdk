package info.sqldev.user;

public class User {
    private long id;
    private java.util.Date created_at;
    private java.util.Date updated_at;
    private int role;
    private String name;
    private String user_name;
    private long state;
    private int sex;
    private String email;
    private String last_ip;
    private java.util.Date last_time;
    private String telno;
    private long expire;
    private String group_name;

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public java.util.Date getCreated_at() {
        return created_at;
    }

    public void setCreated_at(java.util.Date created_at) {
        this.created_at = created_at;
    }

    public java.util.Date getUpdated_at() {
        return updated_at;
    }

    public void setUpdated_at(java.util.Date updated_at) {
        this.updated_at = updated_at;
    }

    public int getRole() {
        return role;
    }

    public void setRole(int role) {
        this.role = role;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getUser_name() {
        return user_name;
    }

    public void setUser_name(String user_name) {
        this.user_name = user_name;
    }

    public long getState() {
        return state;
    }

    public void setState(long state) {
        this.state = state;
    }

    public int getSex() {
        return this.sex;
    }

    public void setSex(int sex) {
        this.sex = sex;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getLast_ip() {
        return last_ip;
    }

    public void setLast_ip(String last_ip) {
        this.last_ip = last_ip;
    }

    public java.util.Date getLast_time() {
        return last_time;
    }

    public void setLast_time(java.util.Date last_time) {
        this.last_time = last_time;
    }

    public String getTelno() {
        return telno;
    }

    public void setTelno(String telno) {
        this.telno = telno;
    }

    public long getExpire() {
        return expire;
    }

    public void setExpire(long expire) {
        this.expire = expire;
    }

    public String getGroup_name() {
        return group_name;
    }

    public void setGroup_name(String group_name) {
        this.group_name = group_name;
    }
}
