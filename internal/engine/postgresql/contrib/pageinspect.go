// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/sqlc-dev/sqlc/internal/sql/ast"
	"github.com/sqlc-dev/sqlc/internal/sql/catalog"
)

var funcsPageinspect = []*catalog.Function{
	{
		Name: "brin_metapage_info",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "brin_page_items",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
			{
				Name: "index_oid",
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "brin_page_type",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "brin_revmap_data",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "tid"},
	},
	{
		Name: "bt_metap",
		Args: []*catalog.Argument{
			{
				Name: "relname",
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "bt_page_items",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "bt_page_items",
		Args: []*catalog.Argument{
			{
				Name: "relname",
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Name: "blkno",
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "bt_page_stats",
		Args: []*catalog.Argument{
			{
				Name: "relname",
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Name: "blkno",
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "fsm_page_contents",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "get_raw_page",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea"},
	},
	{
		Name: "get_raw_page",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea"},
	},
	{
		Name: "gin_leafpage_items",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "gin_metapage_info",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "gin_page_opaque_info",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "gist_page_items",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
			{
				Name: "index_oid",
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "gist_page_items_bytea",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "gist_page_opaque_info",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "hash_bitmap_info",
		Args: []*catalog.Argument{
			{
				Name: "index_oid",
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name: "blkno",
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "hash_metapage_info",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "hash_page_items",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "hash_page_stats",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "hash_page_type",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "heap_page_item_attrs",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
			{
				Name: "rel_oid",
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "heap_page_item_attrs",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
			{
				Name: "rel_oid",
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name: "do_detoast",
				Type: &ast.TypeName{Name: "boolean"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "heap_page_items",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "heap_tuple_infomask_flags",
		Args: []*catalog.Argument{
			{
				Name: "t_infomask",
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Name: "t_infomask2",
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "page_checksum",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
			{
				Name: "blkno",
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "smallint"},
	},
	{
		Name: "page_header",
		Args: []*catalog.Argument{
			{
				Name: "page",
				Type: &ast.TypeName{Name: "bytea"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "tuple_data_split",
		Args: []*catalog.Argument{
			{
				Name: "rel_oid",
				Type: &ast.TypeName{Name: "oid"},
			},
			{
				Name: "t_data",
				Type: &ast.TypeName{Name: "bytea"},
			},
			{
				Name: "t_infomask",
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Name: "t_infomask2",
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Name: "t_bits",
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea[]"},
	},
	{
		Name: "tuple_data_split",
		Args: []*catalog.Argument{
			{
				Name: "rel_oid",
				Type: &ast.TypeName{Name: "oid"},
			},
			{
				Name: "t_data",
				Type: &ast.TypeName{Name: "bytea"},
			},
			{
				Name: "t_infomask",
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Name: "t_infomask2",
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Name: "t_bits",
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Name: "do_detoast",
				Type: &ast.TypeName{Name: "boolean"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea[]"},
	},
}

func Pageinspect() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsPageinspect
	return s
}
