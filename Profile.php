<?php

namespace App;

use Illuminate\Database\Eloquent\Table;


class Profile extends Table
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $table = 'profile'

    protected $fillable = [
        'fullname','age',
    ];
}
