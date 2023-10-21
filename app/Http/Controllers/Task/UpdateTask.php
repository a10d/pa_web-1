<?php

namespace App\Http\Controllers\Task;

use App\Foundation\Http\Controller;
use App\Models\Task;

class UpdateTask extends Controller
{
    public function __invoke(Task $task)
    {
        $this->authorize('update', [$task]);
    }

}
