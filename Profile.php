<?php

namespace App;

use Illuminate\Database\Eloquent\Model;


class Profile extends Model
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
