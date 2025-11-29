#include <php.h>
#include <Zend/zend_API.h>
#include <Zend/zend_hash.h>
#include <Zend/zend_types.h>
#include <stddef.h>

#include "sconcur.h"
#include "sconcur_arginfo.h"
#include "_cgo_export.h"


PHP_MINIT_FUNCTION(sconcur) {
    
    return SUCCESS;
}

zend_module_entry sconcur_module_entry = {STANDARD_MODULE_HEADER,
                                         "sconcur",
                                         ext_functions,             /* Functions */
                                         PHP_MINIT(sconcur),  /* MINIT */
                                         NULL,                      /* MSHUTDOWN */
                                         NULL,                      /* RINIT */
                                         NULL,                      /* RSHUTDOWN */
                                         NULL,                      /* MINFO */
                                         "1.0.0",                   /* Version */
                                         STANDARD_MODULE_PROPERTIES};
PHP_FUNCTION(SConcur_Extension_ping)
{
    zend_string *string = NULL;
    ZEND_PARSE_PARAMETERS_START(1, 1)
        Z_PARAM_STR(string)
    ZEND_PARSE_PARAMETERS_END();
    zend_string *result = ping(string);
    if (result) {
        RETURN_STR(result);
    }

	RETURN_EMPTY_STRING();
}

