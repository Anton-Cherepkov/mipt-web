function addTask(list, taskText)
{
	var newTask = list.append('<li></li>').find('li').last();
	newTask.append('<span>' + taskText + '</span><br>');
	newTask.append('<input type="button" id="remove_task" value="Удалить">');
}

$(function()
{
	var root = $('#root');
	
	root.append('Новое задание: <input type="input" id="add_task_input"><br>');
	root.append('<input type="button" id="add_task" value="Добавить"><br><br>');
	
	root.append('Список заданий:<br>');
	var list = root.append('<ul></ul>').find('ul').last();
	addTask(list, 'Сделать задание #3 по web-программированию');
	
	
	$('#add_task').click(function()
	{
		taskText = $('#add_task_input').val();
		if (/\S/.test(taskText))
		{
			addTask(list, taskText);
		}
		else
		{
			alert('Задание не может быть пустым!');
		}
	});
	
	root.on('click', '#remove_task', function()
	{
		$(this).parent().remove();
	});
});
