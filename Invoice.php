<?php

namespace App;

use Illuminate\Database\Eloquent\Table;


class Invoice extends Table
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $table = 'invoice'

    protected $fillable = [
        'id',
        'ref_no',
        'invoice_date',
        'po_no',
        'quote_id',
        'subtotal',
        'totaltax',
        'delivery_fee',
        'discount',
        'grandtotal',
        'delivery_date',
        'internal_note',
        'invoice_note',
        'invoice_status',
        'delivery_type',
        'delivery_status',
        'created_by',
        'modified_by',
        'deleted_by',
        'deleted_at',
        'created_at',
        'updated_at',
        'promo_id',
        'promo_code',
        
    ];
}
