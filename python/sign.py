import hashlib
import time


def get_signature(params, app_secret):
    sorted_params = sorted(params.items(), key=lambda x: x[0])
    base_string = ''
    for k, v in sorted_params:
        if k != '' and k != 'sign' and k != 'key' and v != '':
            base_string += k + '=' + str(v) + '&'
    if base_string != '':
        base_string = base_string[:-1] + app_secret
    return hashlib.md5(base_string.encode('utf-8')).hexdigest()


if __name__ == '__main__':
    params = {
        'app_key': '0f4199cb4edafbbebf5f1d4cf66c3a1d',
        'current_time': str(int(time.time() * 1000)),
        'role': '0',
        'source': '',
        'page_no': '0',
        'page_size': '10',
    }
    app_secret = '4949ad7f-8d0e-4569-babd-820bb30fc18d'
    print(get_signature(params, app_secret))
