#ifdef HAVE_CONFIG_H
#include "config.h"
#endif

#include <php.h>
#include <stdlib.h>
#include "_cgo_export.h"

// Описание аргументов: 1 обязательный аргумент типа string
ZEND_BEGIN_ARG_INFO_EX(arginfo_ping, 0, 0, 1)
    ZEND_ARG_TYPE_INFO(0, name, IS_STRING, 0)
ZEND_END_ARG_INFO()

// Реализация функции
PHP_FUNCTION(ping) {
    char *name = NULL;
    size_t name_len;

    // "s" означает строку. PHP сам заполнит указатель name и длину name_len
    if (zend_parse_parameters(ZEND_NUM_ARGS(), "s", &name, &name_len) == FAILURE) {
        RETURN_THROWS();
    }

    // Вызов Go
    char *response = ping(name);

    // Возврат значения в PHP (копирование)
    RETVAL_STRING(response);

    // Очистка памяти Go
    free(response);
}

// Регистрация функции с неймспейсом SConcur\Extension
static const zend_function_entry sconcur_functions[] = {
    ZEND_NS_FE("SConcur\\Extension", ping, arginfo_ping)
    PHP_FE_END
};

// Модуль
zend_module_entry sconcur_module_entry = {
    STANDARD_MODULE_HEADER,
    "sconcur",
    sconcur_functions,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    "0.1",
    STANDARD_MODULE_PROPERTIES
};

// Точка входа (без ifdef, чтобы работало при go build)
ZEND_GET_MODULE(sconcur)