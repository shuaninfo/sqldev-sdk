import sqldev

# 为Sqldev类编写python版本的测试
# 常量从下面这段go代码拿
# const (
# 	ENDPOINT   = `http://127.0.0.1:8080/openapi/v1`
# 	APP_KEY    = `a9967fcfa669c40c5df7db31a33f70d4`
# 	APP_SECRET = `5414aabf-db54-4f45-b7e1-1d99aa10e47c`
# )

s = sqldev.Sqldev('http://127.0.0.1:8080/openapi/v1', 'a9967fcfa669c40c5df7db31a33f70d4',
                  '5414aabf-db54-4f45-b7e1-1d99aa10e47c')


def test_query():
    project = 1
    iid = '03af2e960d8736f7724484c263b257f4'
    sql = 'select 1'
    owner = 'testdb'
    schema = '*'
    page_index = 0
    page_size = 10
    need_total = True
    print(s.query(project, iid, sql, owner, schema, page_index, page_size, need_total))


def test_get_user():
    id = 0
    print(s.get_user(id))


def test_get_user_page():
    form = sqldev.UserPageForm()
    form.page_no = 0
    form.page_size = 10
    form.username = ''
    form.nickname = ''
    form.state = ''
    print(s.get_user_page(form))


def test_add_user():
    form = sqldev.UserAddForm()
    form.username = 'test'
    form.nickname = 'test'
    form.password = 'test'
    form.state = '0'
    print(s.add_user(form))


def test_upd_user():
    form = sqldev.UserUpdForm()
    form.id = 1
    form.username = 'test'
    form.nickname = 'test'
    form.password = 'test'
    form.state = '0'
    print(s.upd_user(form))


def test_remove_user():
    id = 1
    print(s.remove_user(id))


def test_state_user():
    form = sqldev.UserStateForm()
    form.id = 1
    form.state = '0'
    print(s.state_user(form))


if __name__ == '__main__':
    test_get_user()
